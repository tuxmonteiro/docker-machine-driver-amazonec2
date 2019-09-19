/*
 * Copyright (c) 2019 Globo.com
 * All rights reserved.
 *
 * This source is subject to the Apache License, Version 2.0.
 * Please see the LICENSE file for more information.
 *
 * Authors: See AUTHORS file
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/tuxmonteiro/docker-machine-driver-amazonec2/amazonec2"
)

func main() {
	plugin.RegisterDriver(amazonec2.NewDriver("", ""))
}
