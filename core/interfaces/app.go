/*
	This package provides interface for app(s), including extern OS native apps, snap apps and flatpak apps
*/

package interfaces

import "github.com/eXtern-OS/common/core/types"

// App has to have Export() and IsPaid()
type App interface {
	Export() types.ExportedApp
	IsPaid() bool
}
