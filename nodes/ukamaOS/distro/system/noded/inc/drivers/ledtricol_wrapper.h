/**
 * Copyright (c) 2021-present, Ukama Inc.
 * All rights reserved.
 *
 * This source code is licensed under the XXX-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#ifndef DRIVERS_LEDTRICOL_WRAPPER_H_
#define DRIVERS_LEDTRICOL_WRAPPER_H_

#include "device.h"

int led_tricol_wrapper_init ();
int led_tricol_wrapper_registration(Device* p_dev);
int led_tricol_wrapper_read_properties(DevObj* obj, void* prop, uint16_t* count);
int led_tricol_wrapper_configure(void* p_dev, void* prop, void* data );
int led_tricol_wrapper_read(void* p_dev, void* prop, void* data);
int led_tricol_wrapper_write(void* p_dev, void* prop, void* data);
int led_tricol_wrapper_enable(void* p_dev, void* prop, void* data);
int led_tricol_wrapper_disable(void* p_dev, void* prop, void* data);

#endif /*DRIVERS_LEDTRICOL_WRAPPER_H_*/
