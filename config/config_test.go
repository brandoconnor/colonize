package config_test

import (
	. "github.com/craigmonson/colonize/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/oleiade/reflections"
)

var _ = Describe("Config/Config", func() {
	basename := "base"
	tmplPath := "foo/bar"
	environment := "dev"

	Describe("LoadConfig", func() {
		Context("given a complete config file", func() {
			cfgPath := "../test/.colonize.yaml"
			rootPath := "../test"
			origin := "../test/foo/bar"
			confIn := LoadConfigInput{
				Environment: environment,
				OriginPath:  origin,
				CfgPath:     cfgPath,
				TmplName:    basename,
				TmplPath:    tmplPath,
				RootPath:    rootPath,
			}
			conf, err := LoadConfig(&confIn)

			It("should not return an error", func() {
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("should return the proper type", func() {
				testConf := ColonizeConfig{}
				Ω(*conf).To(BeAssignableToTypeOf(testConf))
			})

			attributes := map[string]string{
				"Environment": environment,
				"OriginPath":  origin,
				"TmplName":    basename,
				"TmplPath":    tmplPath,
				"CfgPath":     cfgPath,
				"RootPath":    rootPath,

				"Templates_Dir":               "env",
				"Environments_Dir":            "env",
				"Base_Environment_Ext":        "default",
				"Autogenerate_Comment":        "This file generated by colonizer.",
				"Combined_Vals_File":          "_combined.tfvars",
				"Combined_Vars_File":          "_combined_variables.tf",
				"Combined_Tf_File":            "_combined.tf",
				"Combined_Derived_Vars_File":  "_combined_derived.tf",
				"Combined_Derived_Vals_File":  "_combined_derived.tfvars",
				"Combined_Remote_Config_File": "_remote_setup.sh",
				"Remote_Config_File":          "remote_setup.sh",
				"Derived_File":                "derived.tfvars",
				"Vals_File_Env_Post_String":   ".tfvars",
			}

			for k := range attributes {
				var key = k

				It(k+" should be the correct value", func() {
					var val = attributes[key]
					confVal, _ := reflections.GetField(conf, key)
					Ω(confVal).To(Equal(val))
				})
			}

			Context("Generated attributes", func() {
				It("should set TmplRelPaths", func() {
					Ω(conf.TmplRelPaths).To(Equal([]string{"foo", "foo/bar"}))
				})

				It("should set WalkablePaths", func() {
					res := []string{
						"../test",
						"../test/foo",
						"../test/foo/bar",
					}
					Ω(conf.WalkablePaths).To(Equal(res))
				})

				It("should set WalkableValPaths", func() {
					res := []string{
						"../test/env/dev.tfvars",
						"../test/foo/env/dev.tfvars",
						"../test/foo/bar/env/dev.tfvars",
					}
					Ω(conf.WalkableValPaths).To(Equal(res))
				})

				It("should set CombinedValsFilePath", func() {
					expected := "../test/foo/bar/_combined.tfvars"
					Ω(conf.CombinedValsFilePath).To(Equal(expected))
				})

				It("should set CombinedVarsFilePath", func() {
					expected := "../test/foo/bar/_combined_variables.tf"
					Ω(conf.CombinedVarsFilePath).To(Equal(expected))
				})

				It("should set WalkableTfPaths", func() {
					res := []string{
						"../test/env",
						"../test/foo/env",
						"../test/foo/bar/env",
					}
					Ω(conf.WalkableTfPaths).To(Equal(res))
				})

				It("should set CombinedTfFilePath", func() {
					expected := "../test/foo/bar/_combined.tf"
					Ω(conf.CombinedTfFilePath).To(Equal(expected))
				})

				It("should set WalkableDerivedPaths", func() {
					res := []string{
						"../test/env/derived.tfvars",
						"../test/foo/env/derived.tfvars",
						"../test/foo/bar/env/derived.tfvars",
					}
					Ω(conf.WalkableDerivedPaths).To(Equal(res))
				})

				It("should set CombinedDerivedVarFilePath", func() {
					expected := "../test/foo/bar/_combined_derived.tf"
					Ω(conf.CombinedDerivedVarsFilePath).To(Equal(expected))
				})

				It("should set CombinedDerivedValFilePath", func() {
					expected := "../test/foo/bar/_combined_derived.tfvars"
					Ω(conf.CombinedDerivedValsFilePath).To(Equal(expected))
				})

				It("should set CombinedRemoteFilePath", func() {
					expected := "../test/foo/bar/_remote_setup.sh"
					Ω(conf.CombinedRemoteFilePath).To(Equal(expected))
				})

				It("should set RemoteFilePath", func() {
					expected := "../test/env/remote_setup.sh"
					Ω(conf.RemoteFilePath).To(Equal(expected))
				})
			})
		})
	})

	Describe("LoadConfigInTree", func() {
		Context("given a path (tree) to search for the env", func() {
			path := "../test"
			conf, err := LoadConfigInTree(path, environment)

			It("should not return an error", func() {
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("should return the proper type", func() {
				testConf := ColonizeConfig{}
				Ω(*conf).To(BeAssignableToTypeOf(testConf))
			})
		})
	})

	Describe("GetEnvValPath", func() {
		It("should return the environment file path for the env", func() {
			c, _ := LoadConfigInTree("../test/vpc", environment)
			Ω(c.GetEnvValPath()).To(Equal("env/dev.tfvars"))
		})
	})

	Describe("GetEnvTfPath", func() {
		It("should return the environment file path for the env", func() {
			c, _ := LoadConfigInTree("../test/vpc", environment)
			Ω(c.GetEnvTfPath()).To(Equal("env"))
		})
	})

	Describe("GetEnvDerivedPath", func() {
		It("should return the environment file path for the env", func() {
			c, _ := LoadConfigInTree("../test/vpc", environment)
			Ω(c.GetEnvDerivedPath()).To(Equal("env/derived.tfvars"))
		})
	})
})
