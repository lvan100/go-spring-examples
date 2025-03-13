/*
 * Copyright 2025 The Go-Spring Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package app

import (
	_ "simple-server/app/bootstrap"
	_ "simple-server/app/common/grpcsvr"
	_ "simple-server/app/common/handlers/log"
	_ "simple-server/app/common/httpsvr"
	_ "simple-server/app/common/thriftsvr"
	_ "simple-server/app/grpc-controller"
	_ "simple-server/app/http-controller"
	_ "simple-server/app/thrift-controller"
)
