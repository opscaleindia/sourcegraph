package shared

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const nginxConf = `# This config was generated by Sourcegraph.
# You can adjust the configuration to add additional TLS or HTTP features.
# Read more at https://docs.sourcegraph.com/admin/nginx

error_log stderr;
pid /var/run/nginx.pid;

events {
}

http {
	server_tokens off;

	# We can upload large extensions
	client_max_body_size 150M;

	# Don't timeout websockets quickly. Default is 60s. This is the timeout
	# between reads/writes, not the full session timeout.
	proxy_send_timeout 1h;
	proxy_read_timeout 1h;

	access_log off;
	upstream backend {
		server localhost:8080;
	}

	server {
		listen 7080;
		location / {
			proxy_pass http://backend;
			proxy_set_header Host $http_host;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
			proxy_set_header X-Forwarded-Proto $scheme;
		}
	}
}
`

// nginxProcFile will return a procfile entry for nginx, as well as setup
// configuration for it.
func nginxProcFile() (string, error) {
	// Check we can read the config
	path := filepath.Join(os.Getenv("CONFIG_DIR"), "nginx.conf")
	_, err := ioutil.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return "", errors.Wrapf(err, "failed to read nginx configuration at %s", path)
	}

	// Does not exist
	if err != nil {
		err = ioutil.WriteFile(path, []byte(nginxConf), 0600)
		if err != nil {
			return "", errors.Wrapf(err, "failed to generate nginx configuration to %s", path)
		}
	}

	// This is set for the informational message we show once sourcegraph
	// frontend starts. This is so we can advertise the nginx address, rather
	// than the frontend address.
	SetDefaultEnv("SRC_NGINX_HTTP_ADDR", ":7080")

	return fmt.Sprintf(`nginx: nginx -p . -g 'daemon off;' -c %s | grep -v 'could not open error log file'`, path), nil
}
