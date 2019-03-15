package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ConfigAction struct {
	ConfFile string
	ConfType string
	Cmd      *cobra.Command
}

var Config = NewConfigAction()

func init() {

	CodeGen.Cmd.AddCommand(Config.Cmd)
}

func NewConfigAction() *ConfigAction {

	var action = &ConfigAction{}
	command := &cobra.Command{
		Use:   "config",
		Short: "config code generator",
		Long:  "tool for config code generator",
		Run: func(cmd *cobra.Command, args []string) {

			if file, err := os.Stat(action.ConfFile); err != nil || file.IsDir() {
				fmt.Println("config file does not exists or is dir")
				os.Exit(1)
			}
			filename := filepath.Base(action.ConfFile)
			if action.ConfType == "" {
				if filepath.Ext(action.ConfFile) != "" {
					action.ConfType = filepath.Ext(action.ConfFile)
				} else {
					fmt.Println("config type is unkonw")
					os.Exit(1)
				}
			}

			fmt.Println("process ")
			configName := filename[0:strings.Index(filename, action.ConfType)]
			viper.SetConfigName(configName)
			viper.SetConfigType(action.ConfType[1:])
			viper.AddConfigPath(filepath.Dir(action.ConfFile))
			if err := viper.ReadInConfig(); err != nil {
				fmt.Printf("config read exception %s \n", err)
				os.Exit(1)
			}

			buildStruct2(map[string]interface{}{firstToUpper(configName) + "Config": viper.AllSettings()})

		},
	}

	command.Flags().StringVarP(&action.ConfFile, "config-file", "c", "", "config file")
	command.MarkFlagRequired("config-file")

	command.Flags().StringVarP(&action.ConfType, "config-type", "t", "", "config file type : json,yaml")

	action.Cmd = command
	return action
}

func firstToUpper(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}

func buildStruct(vmap map[string]interface{}) {

	if len(vmap) == 0 {
		return
	}

	stack := make(map[string]interface{})
	for k, v := range vmap {
		fmt.Println(k, v)
		switch v.(type) {
		case map[string]interface{}:
			fieldName := firstToUpper(k)
			structName := fieldName + "Config"
			fmt.Printf("%s %s\n", fieldName, structName)
			stack[structName] = v
		case string:
			fmt.Printf("%s %s\n", k, "string")
		case int:
			fmt.Printf("%s %s\n", k, "int")
		default:

		}
	}
	// fmt.Println("}")
	if len(stack) > 0 {
		fmt.Println(len(stack))
		fmt.Println(stack)
		buildStruct2(stack)
	}
}

func buildStruct2(vmap map[string]interface{}) {

	if len(vmap) == 0 {
		return
	}

	stack := make(map[string]interface{})
	for ks, vs := range vmap {
		// fmt.Println(ks, vs)
		fmt.Printf("type %s struct {\n", ks)
		for k, v := range vs.(map[string]interface{}) {
			switch v.(type) {
			case map[string]interface{}:
				fieldName := firstToUpper(k)
				structName := fieldName + "Config"
				fmt.Printf("%s %s\n", fieldName, structName)
				stack[structName] = v
			case string:
				fmt.Printf("%s %s\n", firstToUpper(k), "string")
			case int:
				fmt.Printf("%s %s\n", firstToUpper(k), "int")
			default:

			}
		}
		fmt.Println("}")
	}
	// fmt.Println("}")
	if len(stack) > 0 {
		// fmt.Println(len(stack))
		// fmt.Println(stack)
		buildStruct2(stack)
	}
}

func build(v interface{}) interface{} {

	switch v.(type) {
	case map[string]interface{}:
		submap := make(map[string]interface{})
		ts := v.(map[string]interface{})
		for ks, vs := range ts {
			submap[ks] = build(vs)
		}
		return submap

	case string:
		return "string"
	case int:
		return "int"
	default:
		return nil
	}
}
