package xel

var WinWidth, WinHeight int
var MouseX, MouseY int

type ButtonAction int

type KeyboardKey int
type KeyBehaviorType int

type MouseButton int
type MouseCursor int

const (
	Button_Action_Down ButtonAction = 1 + iota
	Button_Action_Up
)

const (
	Mouse_Button_Left MouseButton = 1 + iota
	Mouse_Button_Right
)

const (
	MouseCursor_Arrow MouseCursor = 1 + iota
	MouseCursor_Horiz_Resize
	MouseCursor_Vert_Resize
	MouseCursor_IBeam
	MouseCursor_CrossHair
	MouseCursor_Hand
)

const (
	Key_Behavior_HOME KeyBehaviorType = 1 + iota
	Key_Behavior_INSERT
	Key_Behavior_PAGE_UP
	Key_Behavior_PAGE_DOWN
	Key_Behavior_DOWN
	Key_Behavior_UP
	Key_Behavior_RIGHT
	Key_Behavior_LEFT
	Key_Behavior_NIL
	Key_Behavior_END
	Key_Behavior_FNUM
	Key_Behavior_DELETE
	Key_Behavior_ENTER
	Key_Behavior_BACKSPACE
	Key_Behavior_TAB
	Key_Behavior_ESCAPE
	Key_Behavior_UNKNOWN
	Key_Behavior_CHAR
)

const (
	Key_APOSTROPHE KeyboardKey = 1 + iota
	Key_SPACE
	Key_COMMA
	Key_MINUS
	Key_PERIOD
	Key_FORWARD_SLASH
	Key_SEMICOLON
	Key_EQUAL
	Key_LEFT_BRACKET
	Key_BACK_SLASH
	Key_RIGHT_BRACKET
	Key_ESCAPE
	Key_ENTER
	Key_TAB
	Key_BACKSPACE
	Key_INSERT
	Key_DELETE

	Key_RIGHT
	Key_LEFT
	Key_UP
	Key_DOWN
	Key_PAGE_UP
	Key_PAGE_DOWN
	Key_HOME
	Key_END

	Key_F1
	Key_F2
	Key_F3
	Key_F4
	Key_F5
	Key_F6
	Key_F7
	Key_F8
	Key_F9
	Key_F10
	Key_F11
	Key_F12

	Key_PAD_0
	Key_PAD_1
	Key_PAD_2
	Key_PAD_3
	Key_PAD_4
	Key_PAD_5
	Key_PAD_6
	Key_PAD_7
	Key_PAD_8
	Key_PAD_9

	Key_PAD_DECIMAL
	Key_PAD_DIVIDE
	Key_PAD_MULTIPLY
	Key_PAD_SUBTRACT
	Key_PAD_ADD
	Key_PAD_ENTER
	Key_PAD_EQUAL

	Key_1
	Key_2
	Key_3
	Key_4
	Key_5
	Key_6
	Key_7
	Key_8
	Key_9
	Key_0

	Key_A
	Key_B
	Key_C
	Key_D
	Key_E
	Key_F
	Key_G
	Key_H
	Key_I
	Key_J
	Key_K
	Key_L
	Key_M
	Key_N
	Key_O
	Key_P
	Key_Q
	Key_R
	Key_S
	Key_T
	Key_U
	Key_V
	Key_W
	Key_X
	Key_Y
	Key_Z
)

