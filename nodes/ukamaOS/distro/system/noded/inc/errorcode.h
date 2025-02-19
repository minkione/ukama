/**
 * Copyright (c) 2021-present, Ukama Inc.
 * All rights reserved.
 *
 * This source code is licensed under the XXX-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#ifndef INC_ERRCODE_H_
#define INC_ERRCODE_H_

#ifdef __cplusplus
extern "C" {
#endif

#include "usys_error.h"

/* Error codes returned by NodeD service */
typedef enum {
    ERR_NODED_WR_FAIL    = (USYS_ERR_APP_BASE_CODE+1),
    ERR_NODED_R_FAIL,
    ERR_NODED_INVALID_JSON_OBJECT,
    ERR_NODED_MW_ERR,
    ERR_NODED_DB_MISSING_FIELD,
    ERR_NODED_INVALID_FIELD,
    ERR_NODED_DISABLED_FIELD,
    ERR_NODED_CRC_FAILURE,
    ERR_NODED_EXCEED_MAX_SIZE,
    ERR_NODED_VALIDATION_FAILURE,
    ERR_NODED_DB_MISSING_INFO,
    ERR_NODED_JSON_PARSER,
    ERR_NODED_DB_MISSING_NODE_INFO,
    ERR_NODED_INVALID_NODE_INFO,
    ERR_NODED_READ_NODE_INFO,
    ERR_NODED_DB_MISSING_NODE_CFG,
    ERR_NODED_INVALID_NODE_CFG,
    ERR_NODED_READ_NODE_CFG,
    ERR_NODED_DB_MISSING_MODULE_INFO,
    ERR_NODED_INVALID_MODULE_INFO,
    ERR_NODED_DB_MISSING_MODULE_CFG,
    ERR_NODED_INVALID_MODULE_CFG,
    ERR_NODED_DB_MISSING_DEVICE_CFG,
    ERR_NODED_INVALID_DEVICE_CFG,
    ERR_NODED_DESERIAL_FAIL,
    ERR_NODED_DB_MISSING_UNIT,
    ERR_NODED_INVALID_UNIT,
    ERR_NODED_DB_MISSING_MODULE,
    ERR_NODED_INVALID_MODULE,
    ERR_NODED_DEV_MISSING,
    ERR_NODED_DEV_PROPERTY_MISSING,
    ERR_NODED_DEV_PROPERTY_IS_NOT_ALERT_TYPE,
    ERR_NODED_DEV_PROPERTY_MARKED_NOT_AVAILABLE,
    ERR_NODED_DEV_PERMISSION_DENIED,
    ERR_NODED_DEV_HWATTR_MISSING,
    ERR_NODED_DEV_DRVR_MISSING,
    ERR_NODED_DEV_API_NOT_SUPPORTED,
    ERR_NODED_DRVR_API_NOT_SUPPORTED,
    ERR_NODED_DRVR_API_NOT_AVAILABLE,
    ERR_NODED_SYSFS_FILE_MISSING,
    ERR_NODED_SYSFS_WRITE_FAILED,
    ERR_NODED_SYSFS_READ_FAILED,
    ERR_NODED_DEV_IRQ_NOT_REG,
    ERR_NODED_THREAD_CREATE_FAIL,
    ERR_NODED_THREAD_CANCEL_FAIL,
    ERR_NODED_MEMORY_EXHAUSTED,
    ERR_NODED_INVALID_POINTER,
    ERR_NODED_LIST_DEL_FAILED,
    ERR_NODED_UNEXPECTED_JSON_OBJECT,
    ERR_NODED_CRT_JSON_SCHEMA,
    ERR_NODED_DB_LNK_MISSING,
    ERR_NODED_DB_MISSING,
    ERR_NODED_JSON_CRETATION_ERR,
    ERR_NODED_JSON_NO_VAL_TO_ENCODE,
    ERR_NODED_JSON_INVALID,
    ERR_NODED_JSON_UNEXPECTED_TAG,
    ERR_NODED_JSON_BAD_REQ,
    ERR_NODED_MAX
} ErrorCode;

#ifdef __cplusplus
}
#endif

#endif /* INC_ERRCODE_H_*/
