package structs

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type EnvironmentHclFile struct {
	Name string
	Path string
}

type GlobalHclFile struct {
}

func (i GlobalHclFile) Write() {
	if err := os.MkdirAll("_base_modules", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	hclFile := hclwrite.NewEmptyFile()
	tfFile, err := os.Create("_base_modules/global.hcl")
	if err != nil {
		fmt.Println(err)
		return
	}
	rootBody := hclFile.Body()
	rootBody.SetAttributeRaw("inputs", hclwrite.TokensForObject([]hclwrite.ObjectAttrTokens{}))
	tfFile.Write(hclFile.Bytes())
}

type BaseModuleHclFile struct {
	Name           string
	TerraformBlock TerraformBlock
}

type AppHclFile struct {
	Path            string
	Environment     string
	Group           string
	App             string
	BaseModule      string
	DependencyChain []string
	Includes        []IncludeBlock
	Dependencies    []DependencyBlock
}

func (i EnvironmentHclFile) Write() {
	if err := os.MkdirAll(i.Path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	hclFile := hclwrite.NewEmptyFile()
	tfFile, err := os.Create(fmt.Sprintf("%s/%s.hcl", i.Path, i.Name))
	if err != nil {
		fmt.Println(err)
		return
	}
	rootBody := hclFile.Body()
	rootBody.SetAttributeRaw("inputs", hclwrite.TokensForObject([]hclwrite.ObjectAttrTokens{}))
	tfFile.Write(hclFile.Bytes())
}

func (i BaseModuleHclFile) Write() {
	var str = fmt.Sprintf("_base_modules/%s", i.Name)
	if err := os.MkdirAll(str, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	hclFile := hclwrite.NewEmptyFile()
	tfFile, err := os.Create(fmt.Sprintf("%s/%s.hcl", str, i.Name))
	if err != nil {
		fmt.Println(err)
		return
	}
	rootBody := hclFile.Body()
	i.TerraformBlock = TerraformBlock{
		Source: i.Name,
	}
	i.TerraformBlock.Body(rootBody)
	rootBody.SetAttributeRaw("inputs", hclwrite.TokensForObject([]hclwrite.ObjectAttrTokens{}))
	tfFile.Write(hclFile.Bytes())
}

func (i GroupHclFile) Write() {
	if err := os.MkdirAll(i.Path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	hclFile := hclwrite.NewEmptyFile()
	tfFile, err := os.Create(fmt.Sprintf("%s/%s.hcl", i.Path, i.Path))
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	rootBody := hclFile.Body()
	remoteStateBlock := RemoteStateBlock{
		Backend:            "azurerm",
		ResourceGroupName:  "test",
		StorageAccountName: "test",
		ContainerName:      "test",
		KeyPrefix:          "eubnpapi",
	}
	remoteStateBlock.Body(rootBody)
	rootBody.SetAttributeRaw("inputs", hclwrite.TokensForObject([]hclwrite.ObjectAttrTokens{}))
	tfFile.Write(hclFile.Bytes())
}

func (i AppHclFile) Write() {
	if err := os.MkdirAll(i.Path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	hclFile := hclwrite.NewEmptyFile()
	tfFile, err := os.Create(fmt.Sprintf("%s/terragrunt.hcl", i.Path))
	if err != nil {
		fmt.Println(err)
		return
	}
	rootBody := hclFile.Body()
	for _, value := range i.Includes {
		value.Body(rootBody)
	}

	for _, value := range i.Dependencies {
		value.Body(rootBody)
	}
	globalInclude := IncludeBlock{
		Name:          "global",
		Path:          "/../../../../_base_modules/global.hcl",
		Expose:        true,
		MergeStrategy: "deep",
	}

	baseModuleInclude := IncludeBlock{
		Name:          i.BaseModule,
		Path:          fmt.Sprintf("/../../../../_base_modules/%s/%s.hcl", i.BaseModule, i.BaseModule),
		Expose:        true,
		MergeStrategy: "deep",
	}

	groupInclude := IncludeBlock{
		Name:          i.Group,
		Path:          fmt.Sprintf("/../../../../%s/%s.hcl", i.Group, i.Group),
		Expose:        true,
		MergeStrategy: "deep",
	}

	environmentInclude := IncludeBlock{
		Name:          i.Environment,
		Path:          fmt.Sprintf("/../../../../%s/%s/%s.hcl", i.Group, i.Environment, i.Environment),
		Expose:        true,
		MergeStrategy: "deep",
	}

	i.Includes = []IncludeBlock{globalInclude, baseModuleInclude, groupInclude, environmentInclude}
	for v := range i.DependencyChain {

		dependencyBlock := DependencyBlock{
			Dependency: i.DependencyChain[v],
		}
		i.Dependencies = append(i.Dependencies, dependencyBlock)

	}

	for v := range i.Includes {

		i.Includes[v].Body(rootBody)

	}

	for v := range i.Dependencies {

		i.Dependencies[v].Body(rootBody)

	}
	rootBody.AppendNewBlock("locals", []string{})
	rootBody.AppendNewline()
	rootBody.SetAttributeRaw("inputs", hclwrite.TokensForObject([]hclwrite.ObjectAttrTokens{}))
	tfFile.Write(hclFile.Bytes())
}

type GroupHclFile struct {
	Path           string
	TerraformBlock TerraformBlock
}

type TerraformBlock struct {
	Source string
}

func (i TerraformBlock) Body(contents *hclwrite.Body) {
	terraform := contents.AppendNewBlock("terraform", []string{})

	tokensSource := hclwrite.Tokens{
		{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenTemplateInterp, Bytes: []byte(`${`)},
		{Type: hclsyntax.TokenIdent, Bytes: []byte(`get_terragrunt_dir`)},
		{Type: hclsyntax.TokenOParen, Bytes: []byte(`(`)},
		{Type: hclsyntax.TokenCParen, Bytes: []byte(`)`)},
		{Type: hclsyntax.TokenTemplateSeqEnd, Bytes: []byte(`}`)},
		{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(fmt.Sprintf(`/../../../../_base_modules/%s`, i.Source))},
		{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
	}
	terraformBody := terraform.Body()
	terraformBody.SetAttributeRaw("source", tokensSource)
	extraArgumentsBlock := terraformBody.AppendNewBlock("extra_arguments", []string{"common_vars"})
	extraArgumentsBody := extraArgumentsBlock.Body()

	tokensCommands := hclwrite.Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte(`get_terraform_commands_that_need_vars`)},
		{Type: hclsyntax.TokenOParen, Bytes: []byte(`(`)},
		{Type: hclsyntax.TokenCParen, Bytes: []byte(`)`)},
	}

	tokensRequiredVarFiles := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrack, Bytes: []byte(`[`)},
		{Type: hclsyntax.TokenCBrack, Bytes: []byte(`]`)},
	}
	extraArgumentsBody.SetAttributeRaw("commands", tokensCommands)
	extraArgumentsBody.SetAttributeRaw("required_var_files", tokensRequiredVarFiles)
	contents.AppendNewline()
	return
}

type IncludeBlock struct {
	Name          string
	Path          string
	Expose        bool
	MergeStrategy string
}

func (i IncludeBlock) Body(contents *hclwrite.Body) {
	groupInclude := contents.AppendNewBlock("include", []string{i.Name})
	groupIncludeBody := groupInclude.Body()
	tokens := hclwrite.Tokens{
		{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenTemplateInterp, Bytes: []byte(`${`)},
		{Type: hclsyntax.TokenIdent, Bytes: []byte(`get_terragrunt_dir`)},
		{Type: hclsyntax.TokenOParen, Bytes: []byte(`(`)},
		{Type: hclsyntax.TokenCParen, Bytes: []byte(`)`)},
		{Type: hclsyntax.TokenTemplateSeqEnd, Bytes: []byte(`}`)},
		{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(i.Path)},
		{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
	}

	groupIncludeBody.SetAttributeRaw("path", tokens)
	groupIncludeBody.SetAttributeValue("expose", cty.BoolVal(i.Expose))
	groupIncludeBody.SetAttributeValue("merge_strategy", cty.StringVal(i.MergeStrategy))
	contents.AppendNewline()
	return
}

type DependencyBlock struct {
	Dependency string
}

func (i DependencyBlock) Body(contents *hclwrite.Body) {
	str := i.Dependency
	replace := "."
	slash := "/"
	underscore := "_"
	n := 2
	result_path := fmt.Sprintf("%s/%s", "../..", strings.Replace(str, replace, slash, n))
	result_name := strings.Replace(str, replace, underscore, n)
	provider := contents.AppendNewBlock("dependency", []string{result_name})
	providerBody := provider.Body()
	providerBody.SetAttributeValue("source_path", cty.StringVal(result_path))
	contents.AppendNewline()
	return
}

type RemoteStateBlock struct {
	Backend            string
	ResourceGroupName  string
	StorageAccountName string
	ContainerName      string
	KeyPrefix          string
}

func (i RemoteStateBlock) Body(contents *hclwrite.Body) {
	remoteStateBlock := contents.AppendNewBlock("remote_state", []string{})
	remoteStateBlockBody := remoteStateBlock.Body()
	remoteStateBlockBody.SetAttributeValue("backend", cty.StringVal(i.Backend))
	tokensGenerate := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrace, Bytes: []byte(`{`)},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		{Type: hclsyntax.TokenTabs, Bytes: []byte("\t")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte(`path`)},
		{Type: hclsyntax.TokenEqual, Bytes: []byte(`=`)},
		{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(`backend.tf`)},
		{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		{Type: hclsyntax.TokenTabs, Bytes: []byte("\t")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte(`if_exists`)},
		{Type: hclsyntax.TokenEqual, Bytes: []byte(`=`)},
		{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(`overwrite`)},
		{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		{Type: hclsyntax.TokenTemplateSeqEnd, Bytes: []byte(`}`)},
	}
	remoteStateBlockBody.SetAttributeRaw("generate", tokensGenerate)
	tokensConfig := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrace, Bytes: []byte(`{`)},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		{Type: hclsyntax.TokenTabs, Bytes: []byte("\t")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte(`resource_group_name`)},
		{Type: hclsyntax.TokenEqual, Bytes: []byte(`=`)},
		{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(i.ResourceGroupName)},
		{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		{Type: hclsyntax.TokenTabs, Bytes: []byte("\t")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte(`storage_account_name`)},
		{Type: hclsyntax.TokenEqual, Bytes: []byte(`=`)},
		{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(i.StorageAccountName)},
		{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		{Type: hclsyntax.TokenTabs, Bytes: []byte("\t")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte(`container_name`)},
		{Type: hclsyntax.TokenEqual, Bytes: []byte(`=`)},
		{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(i.ContainerName)},
		{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		{Type: hclsyntax.TokenTabs, Bytes: []byte("\t")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte(`key`)},
		{Type: hclsyntax.TokenEqual, Bytes: []byte(`=`)},
		{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(i.KeyPrefix)},
		{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(`/`)},
		{Type: hclsyntax.TokenTemplateInterp, Bytes: []byte(`${`)},
		{Type: hclsyntax.TokenIdent, Bytes: []byte(`path_relative_to_include`)},
		{Type: hclsyntax.TokenOParen, Bytes: []byte(`(`)},
		{Type: hclsyntax.TokenCParen, Bytes: []byte(`)`)},
		{Type: hclsyntax.TokenTemplateSeqEnd, Bytes: []byte(`}`)},
		{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(`.tfstate`)},
		{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		{Type: hclsyntax.TokenTabs, Bytes: []byte("\t")},
		{Type: hclsyntax.TokenTemplateSeqEnd, Bytes: []byte(`}`)},
	}
	remoteStateBlockBody.SetAttributeRaw("config", tokensConfig)
	return
}
