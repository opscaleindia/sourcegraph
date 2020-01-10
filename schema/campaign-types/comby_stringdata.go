// Code generated by stringdata. DO NOT EDIT.

package schema

// CombyCampaignTypeSchemaJSON is the content of the file "campaign-types/comby.schema.json".
const CombyCampaignTypeSchemaJSON = `{
  "$id": "comby-spec.json#",
  "$schema": "http://json-schema.org/draft-07/schema#",
  "description": "Schema for comby options",
  "type": "object",
  "properties": {
    "scopeQuery": {
      "type": "string",
      "minLength": 1,
      "description": "Define a scope to narrow down repositories affected by this change. Only GitHub and Bitbucket Server are supported."
    },
    "matchTemplate": {
      "type": "string",
      "minLength": 1,
      "description": "See https://comby.dev/#match-syntax for syntax"
    },
    "rewriteTemplate": {
      "type": "string",
      "description": "See https://comby.dev/#match-syntax for syntax"
    }
  },
  "required": ["scopeQuery", "matchTemplate", "rewriteTemplate"],
  "additionalProperties": false
}
`
