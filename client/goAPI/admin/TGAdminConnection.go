/**
 * Copyright 2018-19 TIBCO Software Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); You may not use this file except
 * in compliance with the License.
 * A copy of the License is included in the distribution package with this file.
 * You also may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF DirectionAny KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * File name: TGAdminConnection.go
 * Created on: Mar 03, 2019
 * Created by: achavan
 * SVN id: $id: $
 *
 */

package admin

import "github.com/TIBCOSoftware/tgdb-client/client/goAPI/types"

type TGAdminConnection interface {
	types.TGConnection
	// CheckpointServer allows the programmatic control to do a checkpoint on server
	CheckpointServer() types.TGError

	// DumpServerStackTrace allows the programmatic control to dump the stack trace on the server console
	DumpServerStackTrace() types.TGError

	// GetAttributeDescriptors gets the list of attribute descriptors
	GetAttributeDescriptors() ([]types.TGAttributeDescriptor, types.TGError)

	// GetConnections gets the list of all socket connections using this connection type
	GetConnections() ([]TGConnectionInfo, types.TGError)

	// GetIndices gets the list of all indices
	GetIndices() ([]TGIndexInfo, types.TGError)

	// GetInfo retrieves the server information (including the server status, listener information,
	// memory information, transaction statistics, cache statistics, database statistics)
	GetInfo() (TGServerInfo, types.TGError)

	// GetUsers gets the list of users
	GetUsers() ([]TGUserInfo, types.TGError)

	// KillConnection allows the programmatic control to stop a particular connection instance
	KillConnection(sessionId int64) types.TGError

	// SetServerLogLevel sets the appropriate log level on server
	SetServerLogLevel(logLevel int, logComponent int64) types.TGError

	// StopServer allows the programmatic-stop of the server execution
	StopServer() types.TGError
}
