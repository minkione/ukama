/**
 * Copyright (c) 2021-present, Ukama Inc.
 * All rights reserved.
 *
 * This source code is licensed under the XXX-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#ifndef CALLBACK_H
#define CALLBACK_H

#include <ulfius.h>

#include "mesh.h"

int callback_webservice(const URequest *request, UResponse *response,
		       void *user_data);
int callback_websocket(const URequest *request, UResponse *response,
		       void *user_data);
int callback_not_allowed(const URequest *request, UResponse *response,
			 void *user_data);
int callback_default_webservice(const URequest *request, UResponse *response,
				void *user_data);
int callback_default_websocket(const URequest *request, UResponse *response,
			       void *user_data);
#endif /* CALLBACK_H */
