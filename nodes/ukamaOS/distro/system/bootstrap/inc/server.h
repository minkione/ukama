/**
 * Copyright (c) 2022-present, Ukama Inc.
 * All rights reserved.
 *
 * This source code is licensed under the XXX-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

/*
 * server.h
 */

#ifndef SERVER_H
#define SERVER_H

#define TRUE  1
#define FALSE 0

#define KEY_NODE         "node"
#define KEY_LOOKING_FOR  "looking_for"
#define VALUE_VALIDATION "validation"

#define MAX_GET_URL_LEN 2048

typedef struct _response {
	char *buffer;
	size_t size;
} Response;

/* Struct to define the server */
typedef struct {

	char *IP;   /* Server's IPv4 for Mesh.d */
	char *cert; /* Cert for connection with Server */
	char *org;  /* Organization this Node belong's */
} ServerInfo;

int register_to_server(char *bootstrapServer, char *uuid, ServerInfo *server);
void free_server_info(ServerInfo *server);
void log_debug_server(ServerInfo *server);
#endif /* SERVER_H */
