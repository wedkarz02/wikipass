package consts

const (
	SecretDir      = "./secret"
	EncryptionFile = "./secret/passwd.aes"
	IVFile         = "./secret/initvec.bin"
	AuthFile       = "./secret/auth.aes"
)

const (
	LogWindowWidth  = 400
	LogWindowHeight = 500
	LogWindowPosX   = 1920/2 - LogWindowWidth/2
	LogWindowPosY   = 1080/2 - LogWindowHeight/2

	AppWindowWidth  = 800
	AppWindowHeight = 600
	AppWindowPosX   = 1920/2 - AppWindowWidth/2
	AppWindowPosY   = 1080/2 - AppWindowHeight/2

	LogoWidth  = 90
	LogoHeight = 90
)
