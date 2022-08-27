package cmd

// IsValidArgs checks length of args from input stdin
func IsValidArgs(args []string)bool  {
    return len(args)>1
}
