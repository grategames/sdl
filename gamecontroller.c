// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "gamecontroller.h"

void extractBind(SDL_GameControllerButtonBind bind, SDL_GameControllerBindType* type, int* v1, int* v2) {
	switch(bind.bindType) {
		case SDL_CONTROLLER_BINDTYPE_BUTTON:
			*v1 = bind.value.button;
			break;
		case SDL_CONTROLLER_BINDTYPE_AXIS:
			*v1 = bind.value.axis;
			break;
		case SDL_CONTROLLER_BINDTYPE_HAT:
			*v1 = bind.value.hat.hat;
			*v2 = bind.value.hat.hat_mask;
			break;
	}
}
