/***********************************************************************************
 *                         This file is part of samurai-render-go
 *                    https://github.com/Samudevv/samurai-render-go
 ***********************************************************************************
 * Copyright (c) 2023 Jonas Pucher
 *
 * This software is provided ‘as-is’, without any express or implied
 * warranty. In no event will the authors be held liable for any damages
 * arising from the use of this software.
 *
 * Permission is granted to anyone to use this software for any purpose,
 * including commercial applications, and to alter it and redistribute it
 * freely, subject to the following restrictions:
 *
 * 1. The origin of this software must not be misrepresented; you must not
 * claim that you wrote the original software. If you use this software
 * in a product, an acknowledgment in the product documentation would be
 * appreciated but is not required.
 *
 * 2. Altered source versions must be plainly marked as such, and must not be
 * misrepresented as being the original software.
 *
 * 3. This notice may not be removed or altered from any source
 * distribution.
 ************************************************************************************/

package samure

/*
#include <linux/input-event-codes.h>
#include <wayland-client-protocol.h>
*/
import "C"

// Keyboard keys
const (
	KeyReserved         = C.KEY_RESERVED
	KeyEsc              = C.KEY_ESC
	Key1                = C.KEY_1
	Key2                = C.KEY_2
	Key3                = C.KEY_3
	Key4                = C.KEY_4
	Key5                = C.KEY_5
	Key6                = C.KEY_6
	Key7                = C.KEY_7
	Key8                = C.KEY_8
	Key9                = C.KEY_9
	Key0                = C.KEY_0
	KeyMinus            = C.KEY_MINUS
	KeyEqual            = C.KEY_EQUAL
	KeyBackspace        = C.KEY_BACKSPACE
	KeyTab              = C.KEY_TAB
	KeyQ                = C.KEY_Q
	KeyW                = C.KEY_W
	KeyE                = C.KEY_E
	KeyR                = C.KEY_R
	KeyT                = C.KEY_T
	KeyY                = C.KEY_Y
	KeyU                = C.KEY_U
	KeyI                = C.KEY_I
	KeyO                = C.KEY_O
	KeyP                = C.KEY_P
	KeyLeftbrace        = C.KEY_LEFTBRACE
	KeyRightbrace       = C.KEY_RIGHTBRACE
	KeyEnter            = C.KEY_ENTER
	KeyLeftctrl         = C.KEY_LEFTCTRL
	KeyA                = C.KEY_A
	KeyS                = C.KEY_S
	KeyD                = C.KEY_D
	KeyF                = C.KEY_F
	KeyG                = C.KEY_G
	KeyH                = C.KEY_H
	KeyJ                = C.KEY_J
	KeyK                = C.KEY_K
	KeyL                = C.KEY_L
	KeySemicolon        = C.KEY_SEMICOLON
	KeyApostrophe       = C.KEY_APOSTROPHE
	KeyGrave            = C.KEY_GRAVE
	KeyLeftshift        = C.KEY_LEFTSHIFT
	KeyBackslash        = C.KEY_BACKSLASH
	KeyZ                = C.KEY_Z
	KeyX                = C.KEY_X
	KeyC                = C.KEY_C
	KeyV                = C.KEY_V
	KeyB                = C.KEY_B
	KeyN                = C.KEY_N
	KeyM                = C.KEY_M
	KeyComma            = C.KEY_COMMA
	KeyDot              = C.KEY_DOT
	KeySlash            = C.KEY_SLASH
	KeyRightshift       = C.KEY_RIGHTSHIFT
	KeyKpasterisk       = C.KEY_KPASTERISK
	KeyLeftalt          = C.KEY_LEFTALT
	KeySpace            = C.KEY_SPACE
	KeyCapslock         = C.KEY_CAPSLOCK
	KeyF1               = C.KEY_F1
	KeyF2               = C.KEY_F2
	KeyF3               = C.KEY_F3
	KeyF4               = C.KEY_F4
	KeyF5               = C.KEY_F5
	KeyF6               = C.KEY_F6
	KeyF7               = C.KEY_F7
	KeyF8               = C.KEY_F8
	KeyF9               = C.KEY_F9
	KeyF10              = C.KEY_F10
	KeyNumlock          = C.KEY_NUMLOCK
	KeyScrolllock       = C.KEY_SCROLLLOCK
	KeyKp7              = C.KEY_KP7
	KeyKp8              = C.KEY_KP8
	KeyKp9              = C.KEY_KP9
	KeyKpminus          = C.KEY_KPMINUS
	KeyKp4              = C.KEY_KP4
	KeyKp5              = C.KEY_KP5
	KeyKp6              = C.KEY_KP6
	KeyKpplus           = C.KEY_KPPLUS
	KeyKp1              = C.KEY_KP1
	KeyKp2              = C.KEY_KP2
	KeyKp3              = C.KEY_KP3
	KeyKp0              = C.KEY_KP0
	KeyKpdot            = C.KEY_KPDOT
	KeyZenkakuhankaku   = C.KEY_ZENKAKUHANKAKU
	Key_102Nd           = C.KEY_102ND
	KeyF11              = C.KEY_F11
	KeyF12              = C.KEY_F12
	KeyRo               = C.KEY_RO
	KeyKatakana         = C.KEY_KATAKANA
	KeyHiragana         = C.KEY_HIRAGANA
	KeyHenkan           = C.KEY_HENKAN
	KeyKatakanahiragana = C.KEY_KATAKANAHIRAGANA
	KeyMuhenkan         = C.KEY_MUHENKAN
	KeyKpjpcomma        = C.KEY_KPJPCOMMA
	KeyKpenter          = C.KEY_KPENTER
	KeyRightctrl        = C.KEY_RIGHTCTRL
	KeyKpslash          = C.KEY_KPSLASH
	KeySysrq            = C.KEY_SYSRQ
	KeyRightalt         = C.KEY_RIGHTALT
	KeyLinefeed         = C.KEY_LINEFEED
	KeyHome             = C.KEY_HOME
	KeyUp               = C.KEY_UP
	KeyPageup           = C.KEY_PAGEUP
	KeyLeft             = C.KEY_LEFT
	KeyRight            = C.KEY_RIGHT
	KeyEnd              = C.KEY_END
	KeyDown             = C.KEY_DOWN
	KeyPagedown         = C.KEY_PAGEDOWN
	KeyInsert           = C.KEY_INSERT
	KeyDelete           = C.KEY_DELETE
	KeyMacro            = C.KEY_MACRO
	KeyMute             = C.KEY_MUTE
	KeyVolumedown       = C.KEY_VOLUMEDOWN
	KeyVolumeup         = C.KEY_VOLUMEUP
	KeyPower            = C.KEY_POWER
	KeyKpequal          = C.KEY_KPEQUAL
	KeyKpplusminus      = C.KEY_KPPLUSMINUS
	KeyPause            = C.KEY_PAUSE
	KeyScale            = C.KEY_SCALE
	KeyKpcomma          = C.KEY_KPCOMMA
	KeyHangeul          = C.KEY_HANGEUL
	KeyHanguel          = C.KEY_HANGUEL
	KeyHanja            = C.KEY_HANJA
	KeyYen              = C.KEY_YEN
	KeyLeftmeta         = C.KEY_LEFTMETA
	KeyRightmeta        = C.KEY_RIGHTMETA
	KeyCompose          = C.KEY_COMPOSE
	KeyStop             = C.KEY_STOP
	KeyAgain            = C.KEY_AGAIN
	KeyProps            = C.KEY_PROPS
	KeyUndo             = C.KEY_UNDO
	KeyFront            = C.KEY_FRONT
	KeyCopy             = C.KEY_COPY
	KeyOpen             = C.KEY_OPEN
	KeyPaste            = C.KEY_PASTE
	KeyFind             = C.KEY_FIND
	KeyCut              = C.KEY_CUT
	KeyHelp             = C.KEY_HELP
	KeyMenu             = C.KEY_MENU
	KeyCalc             = C.KEY_CALC
	KeySetup            = C.KEY_SETUP
	KeySleep            = C.KEY_SLEEP
	KeyWakeup           = C.KEY_WAKEUP
	KeyFile             = C.KEY_FILE
	KeySendfile         = C.KEY_SENDFILE
	KeyDeletefile       = C.KEY_DELETEFILE
	KeyXfer             = C.KEY_XFER
	KeyProg1            = C.KEY_PROG1
	KeyProg2            = C.KEY_PROG2
	KeyWww              = C.KEY_WWW
	KeyMsdos            = C.KEY_MSDOS
	KeyCoffee           = C.KEY_COFFEE
	KeyScreenlock       = C.KEY_SCREENLOCK
	KeyRotateDisplay    = C.KEY_ROTATE_DISPLAY
	KeyDirection        = C.KEY_DIRECTION
	KeyCyclewindows     = C.KEY_CYCLEWINDOWS
	KeyMail             = C.KEY_MAIL
	KeyBookmarks        = C.KEY_BOOKMARKS
	KeyComputer         = C.KEY_COMPUTER
	KeyBack             = C.KEY_BACK
	KeyForward          = C.KEY_FORWARD
	KeyClosecd          = C.KEY_CLOSECD
	KeyEjectcd          = C.KEY_EJECTCD
	KeyEjectclosecd     = C.KEY_EJECTCLOSECD
	KeyNextsong         = C.KEY_NEXTSONG
	KeyPlaypause        = C.KEY_PLAYPAUSE
	KeyPrevioussong     = C.KEY_PREVIOUSSONG
	KeyStopcd           = C.KEY_STOPCD
	KeyRecord           = C.KEY_RECORD
	KeyRewind           = C.KEY_REWIND
	KeyPhone            = C.KEY_PHONE
	KeyIso              = C.KEY_ISO
	KeyConfig           = C.KEY_CONFIG
	KeyHomepage         = C.KEY_HOMEPAGE
	KeyRefresh          = C.KEY_REFRESH
	KeyExit             = C.KEY_EXIT
	KeyMove             = C.KEY_MOVE
	KeyEdit             = C.KEY_EDIT
	KeyScrollup         = C.KEY_SCROLLUP
	KeyScrolldown       = C.KEY_SCROLLDOWN
	KeyKpleftparen      = C.KEY_KPLEFTPAREN
	KeyKprightparen     = C.KEY_KPRIGHTPAREN
	KeyNew              = C.KEY_NEW
	KeyRedo             = C.KEY_REDO
	KeyF13              = C.KEY_F13
	KeyF14              = C.KEY_F14
	KeyF15              = C.KEY_F15
	KeyF16              = C.KEY_F16
	KeyF17              = C.KEY_F17
	KeyF18              = C.KEY_F18
	KeyF19              = C.KEY_F19
	KeyF20              = C.KEY_F20
	KeyF21              = C.KEY_F21
	KeyF22              = C.KEY_F22
	KeyF23              = C.KEY_F23
	KeyF24              = C.KEY_F24
	KeyPlaycd           = C.KEY_PLAYCD
	KeyPausecd          = C.KEY_PAUSECD
	KeyProg3            = C.KEY_PROG3
	KeyProg4            = C.KEY_PROG4
	KeyAllApplications  = C.KEY_ALL_APPLICATIONS
	KeyDashboard        = C.KEY_DASHBOARD
	KeySuspend          = C.KEY_SUSPEND
	KeyClose            = C.KEY_CLOSE
	KeyPlay             = C.KEY_PLAY
	KeyFastforward      = C.KEY_FASTFORWARD
	KeyBassboost        = C.KEY_BASSBOOST
	KeyPrint            = C.KEY_PRINT
	KeyHp               = C.KEY_HP
	KeyCamera           = C.KEY_CAMERA
	KeySound            = C.KEY_SOUND
	KeyQuestion         = C.KEY_QUESTION
	KeyEmail            = C.KEY_EMAIL
	KeyChat             = C.KEY_CHAT
	KeySearch           = C.KEY_SEARCH
	KeyConnect          = C.KEY_CONNECT
	KeyFinance          = C.KEY_FINANCE
	KeySport            = C.KEY_SPORT
	KeyShop             = C.KEY_SHOP
	KeyAlterase         = C.KEY_ALTERASE
	KeyCancel           = C.KEY_CANCEL
	KeyBrightnessdown   = C.KEY_BRIGHTNESSDOWN
	KeyBrightnessup     = C.KEY_BRIGHTNESSUP
	KeyMedia            = C.KEY_MEDIA
	KeySwitchvideomode  = C.KEY_SWITCHVIDEOMODE
	KeyKbdillumtoggle   = C.KEY_KBDILLUMTOGGLE
	KeyKbdillumdown     = C.KEY_KBDILLUMDOWN
	KeyKbdillumup       = C.KEY_KBDILLUMUP
	KeySend             = C.KEY_SEND
	KeyReply            = C.KEY_REPLY
	KeyForwardmail      = C.KEY_FORWARDMAIL
	KeySave             = C.KEY_SAVE
	KeyDocuments        = C.KEY_DOCUMENTS
	KeyBattery          = C.KEY_BATTERY
	KeyBluetooth        = C.KEY_BLUETOOTH
	KeyWlan             = C.KEY_WLAN
	KeyUwb              = C.KEY_UWB
	KeyUnknown          = C.KEY_UNKNOWN
	KeyVideoNext        = C.KEY_VIDEO_NEXT
	KeyVideoPrev        = C.KEY_VIDEO_PREV
	KeyBrightnessCycle  = C.KEY_BRIGHTNESS_CYCLE
	KeyBrightnessAuto   = C.KEY_BRIGHTNESS_AUTO
	KeyBrightnessZero   = C.KEY_BRIGHTNESS_ZERO
	KeyDisplayOff       = C.KEY_DISPLAY_OFF
	KeyWwan             = C.KEY_WWAN
	KeyWimax            = C.KEY_WIMAX
	KeyRfkill           = C.KEY_RFKILL
	KeyMicmute          = C.KEY_MICMUTE
)

// Mouse buttons
const (
	ButtonMisc          = C.BTN_MISC
	Button0             = C.BTN_0
	Button1             = C.BTN_1
	Button2             = C.BTN_2
	Button3             = C.BTN_3
	Button4             = C.BTN_4
	Button5             = C.BTN_5
	Button6             = C.BTN_6
	Button7             = C.BTN_7
	Button8             = C.BTN_8
	Button9             = C.BTN_9
	ButtonMouse         = C.BTN_MOUSE
	ButtonLeft          = C.BTN_LEFT
	ButtonRight         = C.BTN_RIGHT
	ButtonMiddle        = C.BTN_MIDDLE
	ButtonSide          = C.BTN_SIDE
	ButtonExtra         = C.BTN_EXTRA
	ButtonForward       = C.BTN_FORWARD
	ButtonBack          = C.BTN_BACK
	ButtonTask          = C.BTN_TASK
	ButtonJoystick      = C.BTN_JOYSTICK
	ButtonTrigger       = C.BTN_TRIGGER
	ButtonThumb         = C.BTN_THUMB
	ButtonThumb2        = C.BTN_THUMB2
	ButtonTop           = C.BTN_TOP
	ButtonTop2          = C.BTN_TOP2
	ButtonPinkie        = C.BTN_PINKIE
	ButtonBase          = C.BTN_BASE
	ButtonBase2         = C.BTN_BASE2
	ButtonBase3         = C.BTN_BASE3
	ButtonBase4         = C.BTN_BASE4
	ButtonBase5         = C.BTN_BASE5
	ButtonBase6         = C.BTN_BASE6
	ButtonDead          = C.BTN_DEAD
	ButtonGamepad       = C.BTN_GAMEPAD
	ButtonSouth         = C.BTN_SOUTH
	ButtonA             = C.BTN_A
	ButtonEast          = C.BTN_EAST
	ButtonB             = C.BTN_B
	ButtonC             = C.BTN_C
	ButtonNorth         = C.BTN_NORTH
	ButtonX             = C.BTN_X
	ButtonWest          = C.BTN_WEST
	ButtonY             = C.BTN_Y
	ButtonZ             = C.BTN_Z
	ButtonTl            = C.BTN_TL
	ButtonTr            = C.BTN_TR
	ButtonTl2           = C.BTN_TL2
	ButtonTr2           = C.BTN_TR2
	ButtonSelect        = C.BTN_SELECT
	ButtonStart         = C.BTN_START
	ButtonMode          = C.BTN_MODE
	ButtonThumbl        = C.BTN_THUMBL
	ButtonThumbr        = C.BTN_THUMBR
	ButtonDigi          = C.BTN_DIGI
	ButtonToolPen       = C.BTN_TOOL_PEN
	ButtonToolRubber    = C.BTN_TOOL_RUBBER
	ButtonToolBrush     = C.BTN_TOOL_BRUSH
	ButtonToolPencil    = C.BTN_TOOL_PENCIL
	ButtonToolAirbrush  = C.BTN_TOOL_AIRBRUSH
	ButtonToolFinger    = C.BTN_TOOL_FINGER
	ButtonToolMouse     = C.BTN_TOOL_MOUSE
	ButtonToolLens      = C.BTN_TOOL_LENS
	ButtonToolQuinttap  = C.BTN_TOOL_QUINTTAP
	ButtonStylus3       = C.BTN_STYLUS3
	ButtonTouch         = C.BTN_TOUCH
	ButtonStylus        = C.BTN_STYLUS
	ButtonStylus2       = C.BTN_STYLUS2
	ButtonToolDoubletap = C.BTN_TOOL_DOUBLETAP
	ButtonToolTripletap = C.BTN_TOOL_TRIPLETAP
	ButtonToolQuadtap   = C.BTN_TOOL_QUADTAP
	ButtonWheel         = C.BTN_WHEEL
	ButtonGearDown      = C.BTN_GEAR_DOWN
	ButtonGearUp        = C.BTN_GEAR_UP
)

// Button and Key states
const (
	StateReleased = C.WL_POINTER_BUTTON_STATE_RELEASED
	StatePressed  = C.WL_POINTER_BUTTON_STATE_PRESSED
)
