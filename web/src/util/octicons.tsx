import Octicon, {
    Checklist,
    Diff,
    GitCommit,
    GitPullRequest,
    OcticonProps,
    Tasklist,
    Zap,
} from '@primer/octicons-react'
import React from 'react'

// The type definitions for Octicon's props erroneously omit className, so patch them to include
// className.
//
// TODO(sqs): Remove this when https://github.com/primer/octicons/pull/271 is merged.
declare module '@primer/octicons-react' {
    export interface OcticonProps {
        className: string
    }
}

const defaultProps: Partial<OcticonProps> = { size: 24 }

export const GitPullRequestIcon: React.FunctionComponent<{ className?: string }> = ({ className = '' }) => (
    <Octicon {...defaultProps} icon={GitPullRequest} className={className} />
)

export const DiffIcon: React.FunctionComponent<{ className?: string }> = ({ className = '' }) => (
    <Octicon {...defaultProps} icon={Diff} className={className} />
)

export const GitCommitIcon: React.FunctionComponent<{ className?: string }> = ({ className = '' }) => (
    <Octicon {...defaultProps} icon={GitCommit} className={className} />
)

export const ChecklistIcon: React.FunctionComponent<{ className?: string }> = ({ className = '' }) => (
    <Octicon {...defaultProps} icon={Checklist} className={className} />
)

export const ActionsIcon: React.FunctionComponent<{ className?: string }> = ({ className = '' }) => (
    <Octicon {...defaultProps} icon={Zap} className={className} />
)

export const TasklistIcon: React.FunctionComponent<{ className?: string }> = ({ className = '' }) => (
    <Octicon {...defaultProps} icon={Tasklist} className={className} />
)