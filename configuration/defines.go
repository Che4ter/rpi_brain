package configuration

//FLAGS
const SOH = 0x42

//IDs
const SWITCH_STATE = 0x01
const PING = 0x10
const RESET = 0x11
const STOP = 0x12

//TYPE
const REQUEST = 0x0
const RESPONSE = 0x1

//STATE IDs
const STATE_START = 0x20
const STATE_DRIVE_STRAIGHT = 0x21
const STATE_DRIVE_STRAIGHT_SLOW = 0x22
const STATE_DRIVE_STRAIGHT_FAST = 0x23
const STATE_DRIVE_CURVE_LEFT = 0x24
const STATE_DRIVE_CURVE_RIGHT = 0x25
const STATE_OBSTACLE_STAIR = 0x26
const STATE_OBSTACLE_CROSSING = 0x27
const STATE_OBSTACLE_CROSSBAR = 0x28
const STATE_OBSTACLE_RAMP = 0x29
const STATE_MODUS_BUTTON = 0x30
const STATE_WAIT = 0x31
const STATE_STOP = 0x32
const STATE_MANUAL_CONTROL = 0x33
