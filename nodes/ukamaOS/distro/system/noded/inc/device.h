/**
 * Copyright (c) 2021-present, Ukama Inc.
 * All rights reserved.
 *
 * This source code is licensed under the XXX-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#ifndef INC_DEVICE_H_
#define INC_DEVICE_H_

#ifdef __cplusplus
extern "C" {
#endif

#include "property.h"

/* Alarm States */
#define ALARM_STATE_NO_ALARM_ACTIVE			0x00
#define ALARM_STATE_LOW_ALARM_ACTIVE		0x01
#define ALARM_STATE_HIGH_ALARM_ACTIVE		0x02
#define ALARM_STATE_CRIT_ALARM_ACTIVE		0x03

/* DEVICE TYPE */
#define DEV_TYPE_NULL                       0x0000
#define DEV_TYPE_TMP					    0x0001
#define DEV_TYPE_PWR					    0x0002
#define DEV_TYPE_GPIO					    0x0003
#define DEV_TYPE_LED					    0x0004
#define DEV_TYPE_ADC					    0x0005
#define DEV_TYPE_ATT					    0x0006
#define DEV_TYPE_EEPROM					    0x0007
#define DEV_TYPE_SW						    0x0008
#define DEV_TYPE_MAX                        0x0009

typedef uint16_t DeviceType;

/* DEVICE CLASS */
#define DEV_CLASS_GPIO  			    	0x0001
#define DEV_CLASS_I2C				    	0x0002
#define DEV_CLASS_SPI				    	0x0003
#define DEV_CLASS_UART					    0x0004
typedef uint16_t DeviceClass;

/* GPIO Direction */
#define DEV_GPIO_INPUT                      0
#define DEV_GPIO_OUTPUT                     0

typedef struct  __attribute__((__packed__)) {
    uint8_t bus;
    uint16_t add;
} DevI2cCfg;

typedef struct  __attribute__((__packed__)) {
    uint16_t gpioNum;
    uint8_t direction;
} DevGpioCfg;

typedef struct  __attribute__((__packed__)) {
    DevGpioCfg cs;
    uint8_t bus;
} DevSpiCfg;

typedef struct  __attribute__((__packed__)) {
    uint16_t uartNo;
    /*char tty[64];*/ //this could be sysfs
} DevUartCfg;

typedef struct  __attribute__((__packed__)) {
    void* device;
    void* objAttr;
} DeviceAttr;

#if 0
typedef struct  __attribute__((__packed__)) {
    DeviceType type;
    char name[24];
    char disc[24];
    char mod_UUID[24];
    char sysfile[64];
    void* cfg;
} Device;
#endif

typedef struct  __attribute__((__packed__)) {
    char name[24];
    char desc[24];
    char modUuid[24];
    DeviceType type;
} DevObj;

typedef struct {
    char name[24];
    const void* opsTable;
} DevOpsMap;

//TODO: Not required can be cleanup. Property read is good enough.*/
typedef struct {
    int property;
    char sysFname[64];
}SYSFSMap;

typedef struct {
    uint8_t   alertState;
    int   pidx;
    void* sValue;
} AlertCallBackData;

typedef void (*CallBackFxn)(DevObj *obj, AlertCallBackData** acbdata, int* count);
typedef void (*SensorCallbackFxn)(void* cfg);

typedef struct  __attribute__((__packed__)) {
    DevObj obj;
    char sysFile[64];
    const void* devOps;
    void* hwAttr;
    CallBackFxn devCb;
    Property *prop;
} Device;

#define COMPARE_DEV_OBJ(obj1 , obj2)		( !(strcmp(obj1.name, obj2.name)) && \
                !(strcmp(obj1.disc, obj2.disc)) && \
                !(strcmp(obj1.mod_UUID, obj2.mod_UUID)) ) ? (1):(0) \

#define SIZE_OF_DEVICE_CFG(size , type)			{ size = 0 ; \
                switch (type) { \
                    case DEV_CLASS_GPIO: { \
                        size = sizeof(DevGpioCfg); \
                        break; \
                    } \
                    case DEV_CLASS_I2C: { \
                        size = sizeof(DevI2cCfg); \
                        break; \
                    } \
                    case DEV_CLASS_SPI: { \
                        size = sizeof(DevSpiCfg); \
                        break; \
                    } \
                    case DEV_CLASS_UART: { \
                        size = sizeof(DevUartCfg); \
                        break; \
                    } \
                    default: { \
                        size = 0; \
                    } \
                    \
                }  }\

#ifdef __cplusplus
}
#endif

#endif /* INC_DEVICE_H_ */
