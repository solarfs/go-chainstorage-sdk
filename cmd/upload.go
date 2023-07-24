package main

import (
	"errors"
	chainstoragesdk "github.com/paradeum-team/chainstorage-sdk"
	"github.com/paradeum-team/chainstorage-sdk/model"
	"github.com/spf13/cobra"
	"github.com/ulule/deepcopier"
	"net/http"
	"os"
	"text/template"
)

func init() {
	uploadCmd.Flags().StringP("Bucket", "b", "", "桶名称")
	uploadCmd.Flags().StringP("DataPath", "p", "", "数据路径")
}

var uploadCmd = &cobra.Command{
	Use:     "upload",
	Short:   "upload",
	Long:    "upload object into bucket",
	Example: "gcscmd upload --Bucket=<bucketname> --DataPath=<datapath>",

	Run: func(cmd *cobra.Command, args []string) {
		uploadRun(cmd, args)
	},
}

func uploadRun(cmd *cobra.Command, args []string) {
	// 桶名称
	bucketName, err := cmd.Flags().GetString("Bucket")
	if err != nil {
		processError("upload", err, args)
	}

	// 数据路径
	dataPath, err := cmd.Flags().GetString("DataPath")
	if err != nil {
		processError("upload", err, args)
	}

	sdk, err := chainstoragesdk.New(&applicationConfig)
	if err != nil {
		processError("upload", err, args)
	}

	file, err := os.Open(dataPath)
	defer file.Close()

	// 上传对象
	response, err := sdk.Upload.UploadDataViaStream(bucketName, file)
	if err != nil {
		processError("upload", err, args)
	}

	//// 上传对象
	//response, err := sdk.Upload.UploadData(bucketName, dataPath)
	//if err != nil {
	//	processError("upload", err, args)
	//}

	uploadRunOutput(cmd, args, response)
}

func uploadRunOutput(cmd *cobra.Command, args []string, resp model.ObjectCreateResponse) {
	respCode := int(resp.Code)
	if respCode != http.StatusOK {
		processError("upload", errors.New(resp.Msg), args)
	}

	carUploadOutput := CarUploadOutput{
		RequestId: resp.RequestId,
		Code:      resp.Code,
		Msg:       resp.Msg,
		Status:    resp.Status,
	}

	err := deepcopier.Copy(&resp.Data).To(&carUploadOutput.Data)
	if err != nil {
		processError("upload", err, args)
	}

	templateContent := `
CID: {{.Data.ObjectCid}}
Name: {{.Data.ObjectName}}
`

	t, err := template.New("uploadTemplate").Parse(templateContent)
	if err != nil {
		processError("upload", err, args)
	}

	err = t.Execute(os.Stdout, carUploadOutput)
	if err != nil {
		processError("upload", err, args)
	}
}
