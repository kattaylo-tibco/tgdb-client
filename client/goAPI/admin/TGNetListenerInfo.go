package admin

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
 * File name: TGNetListenerInfo.go
 * Created on: Apr 06, 2019
 * Created by: achavan
 * SVN id: $id: $
 *
 */

// TGNetListenerInfo allows users to retrieve the Net-Listener information from server
type TGNetListenerInfo interface {
	// GetCurrentConnections returns the count of current connections
	GetCurrentConnections() int
	// GetMaxConnections returns the count of max connections
	GetMaxConnections() int
	// GetListenerName returns the listener name
	GetListenerName() string
	// GetPortNumber returns the port detail of this listener
	GetPortNumber() string
}
