# Chainstorage-SDK

<!-- TOC -->

* [1. Introduction](#1-introduction)
* [2. Installation](#2-installation)
* [3. Getting Started](#3-getting-started)
* [4. Function Reference](#4-function-reference)
    * [4.1 Bucket](#41-bucket)
        * [4.1.1 GetBucketList](#411-getbucketlist)
        * [4.1.2 CreateBucket](#412-createbucket)
        * [4.1.3 EmptyBucket](#413-emptybucket)
        * [4.1.4 RemoveBucket](#414-removebucket)
        * [4.1.5 GetBucketByName](#415-getbucketbyname)
        * [4.1.6 GetUsersQuotaByStorageNetworkCode](#416-getusersquotabystoragenetworkcode)
    * [4.2 Object](#42-object)
        * [4.2.1 GetObjectList](#421-getobjectlist)
        * [4.2.2  RemoveObject](#422--removeobject)
        * [4.2.3  RenameObject](#423--renameobject)
        * [4.2.4 MarkObject](#424-markobject)
        * [4.2.5  IsExistObjectByCid](#425--isexistobjectbycid)
        * [4.2.6  GetObjectByName](#426--getobjectbyname)
    * [4.3 CAR](#43-car)
        * [4.3.1 CreateCarFile](#431-createcarfile)
        * [4.3.2 SplitCarFile](#432-splitcarfile)
        * [4.3.3 ReferenceObject](#433-referenceobject)
        * [4.3.4 UploadCarFile](#434-uploadcarfile)
        * [4.3.5 UploadShardingCarFile](#435-uploadshardingcarfile)
        * [4.3.6 ConfirmShardingCarFiles](#436-confirmshardingcarfiles)
        * [4.3.7 GenerateTempFileName](#437-generatetempfilename)
        * [4.3.8 ParseCarFile](#438-parsecarfile)
        * [4.3.9 SliceBigCarFile](#439-slicebigcarfile)
        * [4.3.10 GenerateShardingCarFiles](#4310-generateshardingcarfiles)
        * [4.3.11 UploadData](#4311-uploaddata)
        * [4.3.12 UploadBigCarFile](#4312-uploadbigcarfile)
        * [4.3.13 UploadCarFileExt](#4313-uploadcarfileext)
        * [4.3.14 UploadShardingCarFileExt](#4314-uploadshardingcarfileext)
        * [4.3.15 ImportCarFileExt](#4315-importcarfileext)
        * [4.3.16 ImportShardingCarFileExt](#4316-importshardingcarfileext)
        * [4.3.17 ExtractCarFile](#4317-extractcarfile)
    * [4.4 Other](#44-other)
        * [4.4.1 GetIpfsVersion](#441-getipfsversion)
        * [4.4.2 GetApiVersion](#442-getapiversion)
* [5. Troubleshooting](#5-troubleshooting)
* [6. FAQ](#6-faq)
* [7. Appendix](#7-appendix)
    * [7.1 Glossary](#71-glossary)
        * [Bucket](#bucket)
        * [BucketPageResponse](#bucketpageresponse)
        * [BucketPage](#bucketpage)
        * [BucketCreateResponse](#bucketcreateresponse)
        * [BucketEmptyResponse](#bucketemptyresponse)
        * [BucketRemoveResponse](#bucketremoveresponse)
        * [Object](#object)
        * [ObjectPageResponse](#objectpageresponse)
        * [ObjectPage](#objectpage)
        * [ObjectCreateResponse](#objectcreateresponse)
        * [ObjectRemoveResponse](#objectremoveresponse)
        * [ObjectRenameResponse](#objectrenameresponse)
        * [ObjectMarkResponse](#objectmarkresponse)
        * [ObjectExistResponse](#objectexistresponse)
        * [ObjectExistCheck](#objectexistcheck)
        * [CarFileUploadReq](#carfileuploadreq)
        * [ShardingCarFileUploadResponse](#shardingcarfileuploadresponse)
        * [CarFileUpload](#carfileupload)
        * [ShardingCarFilesVerifyResponse](#shardingcarfilesverifyresponse)
        * [CarResponse](#carresponse)
        * [RootLink](#rootlink)
    * [7.2 Error Codes](#72-error-codes)

<!-- TOC -->

## 1. Introduction

The SDK provides a convenient way to interact with the Chainstorage API and perform various operations on buckets or
data objects. This document serves as a reference for the SDK interface, method descriptions, parameters, return types,
and examples.

## 2. Installation

To use the SDK in your project, follow these steps:

1. Install Go (minimum version 1.20.2) on your system.
2. Create a new Go module or navigate to your existing project directory.
3. Run the following command to add the SDK as a dependency:
   go get github.com/solarfs/go-chainstorage-sdk

## 3. Getting Started

Before using the SDK, make sure you have the following information:

- API Endpoint: The URL of the API server.
- API Key: The authentication key required to access the API.

To initialize the SDK, use the following code snippet:

```go
package main

import "github.com/solarfs/go-chainstorage-sdk"

func main() {
// Set configuration
config := chainstoragesdk.ApplicationConfig{}

// Initialize the SDK
sdk, err := chainstoragesdk.New(&config)

// Start using the SDK functions
// ...
}
```

## 4. Function Reference

### 4.1 Bucket

#### 4.1.1 GetBucketList

Retrieves a list of buckets based on the specified criteria.

```go
GetBucketList(bucketName string, pageSize, pageIndex int) (model.BucketPageResponse, error)
```

**Parameters**

| Name       | Type   | Description                                |
|------------|--------|--------------------------------------------|
| bucketName | string | The name of the bucket to search for.      |
| pageSize   | int    | The maximum number of buckets to retrieve. |
| pageIndex  | int    | The index of the page to retrieve.         |

**Return Type**

| Type                     | Description          |
|--------------------------|----------------------|
| model.BucketPageResponse | Bucket page response |
| error                    | Error, if any        |

**Usage**

```go
bucketName := "mybucket"
pageSize := 10
pageIndex := 1

bucketPageResponse, err := sdk.Bucket.GetBucketList(bucketName, pageSize, pageIndex)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the bucketPageResponse
fmt.Println("Total Buckets:", bucketPageResponse.Data.Count)
fmt.Println("Page Index:", bucketPageResponse.Data.PageIndex)
fmt.Println("Page Size:", bucketPageResponse.Data.PageSize)
fmt.Println("Buckets:")

for _, bucket := range objectPageResponse.Data.List {
fmt.Println("Bucket ID:", bucket.Id)
fmt.Println("Bucket Name:", bucket.BucketName)
fmt.Println("--------------")
}
```

#### 4.1.2 CreateBucket

Creates a new bucket with the specified details.

```go
CreateBucket(bucketName string, storageNetworkCode, bucketPrincipleCode int) (model.BucketCreateResponse, error)
```

**Parameters**

| Name                | Type   | Description           |
|---------------------|--------|-----------------------|
| bucketName          | string | Bucket name           |
| storageNetworkCode  | int    | Storage network code  |
| bucketPrincipleCode | int    | Bucket principle code |

**Return Type**

| Type                       | Description            |
|----------------------------|------------------------|
| model.BucketCreateResponse | Bucket create response |
| error                      | Error, if any          |

**Usage**

```go
bucketName := "my-new-bucket"
storageNetworkCode := 10001
bucketPrincipleCode := 10001

bucketCreateResponse, err := sdk.Bucket.CreateBucket(bucketName, storageNetworkCode, bucketPrincipleCode)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the bucketCreateResponse
fmt.Println("Bucket ID:", bucketCreateResponse.Data.Id)
fmt.Println("Bucket Name:", bucketCreateResponse.Data.BucketName)
```

#### 4.1.3 EmptyBucket

Empties the contents of a bucket.

```go
EmptyBucket(bucketId int) (model.BucketEmptyResponse, error)
```

**Parameters**

| Name     | Type | Description |
|----------|------|-------------|
| bucketId | int  | Bucket ID   |

**Return Type**

| Type                      | Description           |
|---------------------------|-----------------------|
| model.BucketEmptyResponse | Bucket empty response |
| error                     | Error, if any         |

**Usage**

```go
bucketId := 123

bucketEmptyResponse, err := sdk.Bucket.EmptyBucket(bucketId)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the bucketEmptyResponse
fmt.Println("Bucket Emptied:", bucketEmptyResponse.RequestId)
fmt.Println("Status:", bucketEmptyResponse.Status)
```

#### 4.1.4 RemoveBucket

Removes a bucket.

```go
RemoveBucket(bucketId int, autoEmptyBucketData bool) (model.BucketRemoveResponse, error)
```

**Parameters**

| Name                | Type | Description                                                 |
|---------------------|------|-------------------------------------------------------------|
| bucketId            | int  | Bucket ID                                                   |
| autoEmptyBucketData | bool | Flag to indicate whether to empty the bucket before removal |

**Return Type**

| Type                       | Description            |
|----------------------------|------------------------|
| model.BucketRemoveResponse | Bucket remove response |
| error                      | Error, if any          |

**Usage**

```go
bucketId := 123
autoEmptyBucketData := true

bucketRemoveResponse, err := sdk.Bucket.RemoveBucket(bucketId, autoEmptyBucketData)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the bucketRemoveResponse
fmt.Println("Bucket Emptied:", bucketRemoveResponse.RequestId)
fmt.Println("Status:", bucketRemoveResponse.Status)
```

#### 4.1.5 GetBucketByName

Retrieves a bucket by its name.

```go
GetBucketByName(bucketName string) (model.BucketCreateResponse, error)
```

**Parameters**

| Name       | Type   | Description |
|------------|--------|-------------|
| bucketName | string | Bucket name |

**Return Type**

| Type                       | Description            |
|----------------------------|------------------------|
| model.BucketCreateResponse | Bucket create response |
| error                      | Error, if any          |

**Usage**

```go
bucketName := "my-bucket"

bucketResponse, err := sdk.Bucket.GetBucketByName(bucketName)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the bucketResponse
fmt.Println("Bucket ID:", bucketResponse.Data.Id)
fmt.Println("Bucket Name:", bucketResponse.Data.BucketName)
```

#### 4.1.6 GetUsersQuotaByStorageNetworkCode

Retrieves the quota of users based on the storage network code.

```go
GetUsersQuotaByStorageNetworkCode(storageNetworkCode int) (model.UsersQuotaResponse, error)
```

**Parameters**

| Name               | Type | Description          |
|--------------------|------|----------------------|
| storageNetworkCode | int  | Storage network code |

**Return Type**

| Type                     | Description          |
|--------------------------|----------------------|
| model.UsersQuotaResponse | Users quota response |
| error                    | Error, if any        |

**Usage**

```go
storageNetworkCode := 10001

usersQuotaResponse, err := Bucket.GetUsersQuotaByStorageNetworkCode(storageNetworkCode)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the usersQuotaResponse
fmt.Println("Package Plan ID:", usersQuotaResponse.Data.PackagePlanId)
fmt.Println("Start Time:", usersQuotaResponse.Data.StartTime)
fmt.Println("Expired Time:", usersQuotaResponse.Data.ExpiredTime)
fmt.Println("Details:")

for _, detail := range usersQuotaResponse.Data.Details {
fmt.Println("Constraint ID:", detail.ConstraintId)
fmt.Println("Constraint Name:", detail.ConstraintName)
fmt.Println("Limited Quota:", detail.LimitedQuota)
fmt.Println("Available:", detail.Available)
fmt.Println("--------------")
}
```

### 4.2 Object

#### 4.2.1 GetObjectList

Retrieves a list of objects based on the provided parameters.

```go
GetObjectList(bucketId int, objectItem string, pageSize, pageIndex int) (model.ObjectPageResponse, error)
```

**Parameters**

| Name       | Type   | Description                            |
|------------|--------|----------------------------------------|
| bucketId   | int    | Bucket ID.                             |
| objectItem | string | Object item to search for (optional).  |
| pageSize   | int    | Number of objects per page (optional). |
| pageIndex  | int    | Page index (optional).                 |

**Return Value**

| Type                     | Description                                          |
|--------------------------|------------------------------------------------------|
| model.ObjectPageResponse | Object page response containing the list of objects. |
| error                    | Error, if any.                                       |

**Usage**

```go
bucketId := 123
objectItem := "example"
pageSize := 10
pageIndex := 1

objectPageResponse, err := sdk.Object.GetObjectList(bucketId, objectItem, pageSize, pageIndex)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the objectPageResponse
fmt.Println("Total Objects:", objectPageResponse.Data.Count)
fmt.Println("Page Index:", objectPageResponse.Data.PageIndex)
fmt.Println("Page Size:", objectPageResponse.Data.PageSize)
fmt.Println("Objects:")

for _, obj := range objectPageResponse.Data.List {
fmt.Println("Object ID:", obj.Id)
fmt.Println("Object Name:", obj.ObjectName)
fmt.Println("Object Size:", obj.ObjectSize)
fmt.Println("--------------")
}
```

#### 4.2.2  RemoveObject

Removes multiple objects.

```go
RemoveObject(objectIds []int) (model.ObjectRemoveResponse, error)
```

**Parameters**

| Name      | Type  | Description                    |
|-----------|-------|--------------------------------|
| objectIds | []int | Slice of object IDs to remove. |

**Return Value**

| Type                       | Description             |
|----------------------------|-------------------------|
| model.ObjectRemoveResponse | Object remove response. |
| error                      | Error, if any.          |

**Usage**

```go
objectIds := []int{1, 2, 3}

objectRemoveResponse, err := sdk.Object.RemoveObject(objectIds)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the objectRemoveResponse
fmt.Println("Status:", objectRemoveResponse.Status)
```

#### 4.2.3  RenameObject

Renames an object.

```go
RenameObject(objectId int, objectName string, isOverwrite bool) (model.ObjectRenameResponse, error)
```

**Parameters**

| Name        | Type   | Description                                                                  |
|-------------|--------|------------------------------------------------------------------------------|
| objectId    | int    | Object ID.                                                                   |
| objectName  | string | New name for the object.                                                     |
| isOverwrite | bool   | Flag to indicate whether to overwrite an existing object with the same name. |

**Return Value**

| Type                       | Description             |
|----------------------------|-------------------------|
| model.ObjectRenameResponse | Object rename response. |
| error                      | Error, if any.          |

**Usage**

```go
objectId := 123

objectName := "newObjectName"
isOverwrite := true

objectRenameResponse, err := sdk.Object.RenameObject(objectId, objectName, isOverwrite)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the renameResponse
fmt.Println("Status:", objectRenameResponse.Status)
```

#### 4.2.4 MarkObject

Marks or unmarks an object.

```go
MarkObject(objectId int, isMarked bool) (model.ObjectMarkResponse, error)
```

**Parameters**

| Name     | Type | Description                                            |
|----------|------|--------------------------------------------------------|
| objectId | int  | Object ID.                                             |
| isMarked | bool | Flag to indicate whether to mark or unmark the object. |

**Return Value**

| Type                     | Description           |
|--------------------------|-----------------------|
| model.ObjectMarkResponse | Object mark response. |
| error                    | Error, if any.        |

**Usage**

```go
objectId := 123
isMarked := true

objectMarkResponse, err := sdk.Object.MarkObject(objectId, isMarked)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the objectMarkResponse
fmt.Println("Status:", objectMarkResponse.Status)
```

#### 4.2.5  IsExistObjectByCid

Checks if an object exists based on the object CID.

```go
IsExistObjectByCid(objectCid string) (model.ObjectExistResponse, error)
```

**Parameters**

| Name      | Type   | Description |
|-----------|--------|-------------|
| objectCid | string | Object CID. |

**Return Value**

| Type                      | Description            |
|---------------------------|------------------------|
| model.ObjectExistResponse | Object exist response. |
| error                     | Error, if any.         |

**Usage**

```go
objectCid := "QmX8e9sDjbaaA8hGJkPcWYXzgAPJTjiFsZ3kjsvHE9C97x"

objectExistResponse, err := sdk.Object.IsExistObjectByCid(objectCid)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the existResponse
fmt.Println("Object Exist:", objectExistResponse.Data.IsExist)
```

#### 4.2.6  GetObjectByName

Retrieves an object by its name.

```go
GetObjectByName(bucketId int, objectName string) (model.ObjectCreateResponse, error)
```

**Parameters**

| Name       | Type   | Description  |
|------------|--------|--------------|
| bucketId   | int    | Bucket ID.   |
| objectName | string | Object name. |

**Return Value**

| Type                       | Description             |
|----------------------------|-------------------------|
| model.ObjectCreateResponse | Object create response. |
| error                      | Error, if any.          |

**Usage**

```go
bucketId := 123
objectName := "exampleObject"

objectResponse, err := sdk.Object.GetObjectByName(bucketId, objectName)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the objectResponse
fmt.Println("Object ID:", objectResponse.Data.Id)
fmt.Println("Object Name:", objectResponse.Data.ObjectName)
```

### 4.3 CAR

#### 4.3.1 CreateCarFile

Creates a Car file.

```go
CreateCarFile(dataPath string, fileDestination string) error
```

**Parameters**

| Name            | Type   | Description                        |
|-----------------|--------|------------------------------------|
| dataPath        | string | Path to the data.                  |
| fileDestination | string | Destination path for the Car file. |

**Return Value**

| Type  | Description    |
|-------|----------------|
| error | Error, if any. |

**Usage**

```go
dataPath := "path/to/data"
fileDestination := "path/to/car/file.car"

err := sdk.Car.CreateCarFile(dataPath, fileDestination)
if err != nil {
fmt.Println("Error:", err)
return
}

// Car file created successfully
```

#### 4.3.2 SplitCarFile

Splits a Car file into chunked files.

```go
SplitCarFile(carFilePath string, chunkedFileDestinations *[]string) error
```

**Parameters**

| Name                    | Type      | Description                              |
|-------------------------|-----------|------------------------------------------|
| carFilePath             | string    | Path to the Car file.                    |
| chunkedFileDestinations | *[]string | Destination paths for the chunked files. |

**Return Value**

| Type  | Description    |
|-------|----------------|
| error | Error, if any. |

**Usage**

```go
carFilePath := "path/to/car/file.car"
chunkedFileDestinations := []string{"path/to/chunk1.car", "path/to/chunk2.car"}

err := sdk.Car.SplitCarFile(carFilePath, &chunkedFileDestinations)
if err != nil {
fmt.Println("Error:", err)
return
}

// Car file split successfully into chunked files
```

#### 4.3.3 ReferenceObject

References an object.

```go
ReferenceObject(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error)
```

**Parameters**

| Name | Type                    | Description              |
|------|-------------------------|--------------------------|
| req  | *model.CarFileUploadReq | Car file upload request. |

**Return Value**

| Type                       | Description             |
|----------------------------|-------------------------|
| model.ObjectCreateResponse | Object create response. |
| error                      | Error, if any.          |

**Usage**

```go
carFileUploadReq := &model.CarFileUploadReq{
// Set the required fields in the Car file upload request
}

objectResponse, err := sdk.Car.ReferenceObject(carFileUploadReq)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the objectResponse
fmt.Println("Object CID:", objectResponse.Data.ObjectCid)
fmt.Println("Object Name:", objectResponse.Data.ObjectName)
```

#### 4.3.4 UploadCarFile

Uploads a Car file.

```go
UploadCarFile(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error)
```

**Parameters**

| Name | Type                    | Description              |
|------|-------------------------|--------------------------|
| req  | *model.CarFileUploadReq | Car file upload request. |

**Return Value**

| Type                       | Description             |
|----------------------------|-------------------------|
| model.ObjectCreateResponse | Object create response. |
| error                      | Error, if any.          |

**Usage**

```go
carFileUploadReq := &model.CarFileUploadReq{
// Set the required fields in the Car file upload request
}

objectResponse, err := sdk.Car.UploadCarFile(carFileUploadReq)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the objectResponse
fmt.Println("Object CID:", objectResponse.Data.ObjectCid)
fmt.Println("Object Name:", objectResponse.Data.ObjectName)
```

#### 4.3.5 UploadShardingCarFile

Uploads a sharding Car file.

```go
UploadShardingCarFile(req *model.CarFileUploadReq) (model.ShardingCarFileUploadResponse, error)
```

**Parameters**

| Name | Type                    | Description              |
|------|-------------------------|--------------------------|
| req  | *model.CarFileUploadReq | Car file upload request. |

**Return Value**

| Type                                | Description                        |
|-------------------------------------|------------------------------------|
| model.ShardingCarFileUploadResponse | Sharding Car file upload response. |
| error                               | Error, if any.                     |

**Usage**

```go
carFileUploadReq := &model.CarFileUploadReq{
// Set the required fields in the Car file upload request
}

shardingResponse, err := sdk.Car.UploadShardingCarFile(carFileUploadReq)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the shardingResponse
fmt.Println("Object CID:", objectResponse.Data.ObjectCid)
fmt.Println("Object Name:", objectResponse.Data.ObjectName)
```

#### 4.3.6 ConfirmShardingCarFiles

Confirms sharding Car files.

```go
ConfirmShardingCarFiles(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error)
```

**Parameters**

| Name | Type                    | Description              |
|------|-------------------------|--------------------------|
| req  | *model.CarFileUploadReq | Car file upload request. |

**Return Value**

| Type                       | Description             |
|----------------------------|-------------------------|
| model.ObjectCreateResponse | Object create response. |
| error                      | Error, if any.          |

**Usage**

```go
carFileUploadReq := &model.CarFileUploadReq{
// Set the required fields in the Car file upload request
}

objectResponse, err := sdk.Car.ConfirmShardingCarFiles(carFileUploadReq)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the objectResponse
fmt.Println("Object CID:", objectResponse.Data.ObjectCid)
fmt.Println("Object Name:", objectResponse.Data.ObjectName)
```

#### 4.3.7 GenerateTempFileName

Generates a temporary file name.

```go
GenerateTempFileName(prefix, suffix string) string
```

**Parameters**

| Name   | Type   | Description                         |
|--------|--------|-------------------------------------|
| prefix | string | Prefix for the temporary file name. |
| suffix | string | Suffix for the temporary file name. |

**Return Value**

| Type   | Description          |
|--------|----------------------|
| string | Temporary file name. |

**Usage**

```go
prefix := "temp"
suffix := ".txt"

tempFileName := sdk.Car.GenerateTempFileName(prefix, suffix)
fmt.Println("Temporary file name:", tempFileName)
```

#### 4.3.8 ParseCarFile

Parses a Car file.

```go
ParseCarFile(carFilePath string, rootLink *model.RootLink) error
```

**Parameters**

| Name        | Type            | Description            |
|-------------|-----------------|------------------------|
| carFilePath | string          | Path to the Car file.  |
| rootLink    | *model.RootLink | Root link information. |

**Return Value**

| Type  | Description    |
|-------|----------------|
| error | Error, if any. |

**Usage**

```go
carFilePath := "path/to/car/file.car"
rootLink := &model.RootLink{
// Set the root link information
}

err := sdk.Car.ParseCarFile(carFilePath, rootLink)
if err != nil {
fmt.Println("Error:", err)
return
}

// Car file parsed successfully
```

#### 4.3.9 SliceBigCarFile

Slices a big Car file.

```go
SliceBigCarFile(carFilePath string) error
```

**Parameters**

| Name        | Type   | Description           |
|-------------|--------|-----------------------|
| carFilePath | string | Path to the Car file. |

**Return Value**

| Type  | Description    |
|-------|----------------|
| error | Error, if any. |

**Usage**

```go
carFilePath := "path/to/car/file.car"

err := sdk.Car.SliceBigCarFile(carFilePath)
if err != nil {
fmt.Println("Error:", err)
return
}

// Big Car file sliced successfully
```

#### 4.3.10 GenerateShardingCarFiles

Generates sharding Car files.

```go
GenerateShardingCarFiles(req *model.CarFileUploadReq, shardingCarFileUploadReqs *[]model.CarFileUploadReq) error
```

**Parameters**

| Name                      | Type                      | Description                        |
|---------------------------|---------------------------|------------------------------------|
| req                       | *model.CarFileUploadReq   | Car file upload request.           |
| shardingCarFileUploadReqs | *[]model.CarFileUploadReq | Sharding Car file upload requests. |

**Return Value**

| Type  | Description    |
|-------|----------------|
| error | Error, if any. |

**Usage**

```go
carFileUploadReq := &model.CarFileUploadReq{
// Set the required fields in the Car file upload request
}

shardingCarFileUploadReqs := &[]model.CarFileUploadReq{
// Set the required fields in the sharding Car file upload requests
}

err := sdk.Car.GenerateShardingCarFiles(carFileUploadReq, shardingCarFileUploadReqs)
if err != nil {
fmt.Println("Error:", err)
return
}

// Sharding Car files generated successfully
```

#### 4.3.11 UploadData

Uploads data to a bucket.

```go
UploadData(bucketId int, dataPath string) (model.ObjectCreateResponse, error)
```

**Parameters**

| Name     | Type   | Description       |
|----------|--------|-------------------|
| bucketId | int    | Bucket ID.        |
| dataPath | string | Path to the data. |

**Return Value**

| Type                       | Description             |
|----------------------------|-------------------------|
| model.ObjectCreateResponse | Object create response. |
| error                      | Error, if any.          |

**Usage**

```go
bucketId := 123
dataPath := "path/to/data"

objectResponse, err := sdk.Car.UploadData(bucketId, dataPath)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the objectResponse
fmt.Println("Object CID:", objectResponse.Data.ObjectCid)
fmt.Println("Object Name:", objectResponse.Data.ObjectName)
```

#### 4.3.12 UploadBigCarFile

Uploads a big Car file.

```go
UploadBigCarFile(req *model.CarFileUploadReq) (model.ObjectCreateResponse,

error)
```

**Parameters**

| Name | Type                    | Description              |
|------|-------------------------|--------------------------|
| req  | *model.CarFileUploadReq | Car file upload request. |

**Return Value**

| Type                       | Description             |
|----------------------------|-------------------------|
| model.ObjectCreateResponse | Object create response. |
| error                      | Error, if any.          |

**Usage**

```go
carFileUploadReq := &model.CarFileUploadReq{
// Set the required fields in the Car file upload request
}

objectResponse, err := sdk.Car.UploadBigCarFile(carFileUploadReq)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the objectResponse
fmt.Println("Object CID:", objectResponse.Data.ObjectCid)
fmt.Println("Object Name:", objectResponse.Data.ObjectName)
```

#### 4.3.13 UploadCarFileExt

Uploads a Car file with an external reader.

```go
UploadCarFileExt(req *model.CarFileUploadReq, extReader io.Reader) (model.ObjectCreateResponse, error)
```

**Parameters**

| Name      | Type                    | Description                       |
|-----------|-------------------------|-----------------------------------|
| req       | *model.CarFileUploadReq | Car file upload request.          |
| extReader | io.Reader               | External reader for the Car file. |

**Return Value**

| Type                       | Description             |
|----------------------------|-------------------------|
| model.ObjectCreateResponse | Object create response. |
| error                      | Error, if any.          |

**Usage**

```go
carFileUploadReq := &model.CarFileUploadReq{
// Set the required fields in the Car file upload request
}

file, err := os.Open("path/to/car/file.car")
if err != nil {
fmt.Println("Error:", err)
return
}
defer file.Close()

objectResponse, err := sdk.Car.UploadCarFileExt(carFileUploadReq, file)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the objectResponse
fmt.Println("Object CID:", objectResponse.Data.ObjectCid)
fmt.Println("Object Name:", objectResponse.Data.ObjectName)
```

#### 4.3.14 UploadShardingCarFileExt

Uploads a sharding Car file with an external reader.

```go
UploadShardingCarFileExt(req *model.CarFileUploadReq, extReader io.Reader) (model.ShardingCarFileUploadResponse, error)
```

**Parameters**

| Name      | Type                    | Description                       |
|-----------|-------------------------|-----------------------------------|
| req       | *model.CarFileUploadReq | Car file upload request.          |
| extReader | io.Reader               | External reader for the Car file. |

**Return Value**

| Type                                | Description                        |
|-------------------------------------|------------------------------------|
| model.ShardingCarFileUploadResponse | Sharding Car file upload response. |
| error                               | Error, if any.                     |

**Usage**

```go
carFileUploadReq := &model.CarFileUploadReq{
// Set the required fields in the Car file upload request
}

file, err := os.Open("path/to/car/file.car")
if err != nil {
fmt.Println("Error:", err)
return
}
defer file.Close()

shardingResponse, err := sdk.Car.UploadShardingCarFileExt(carFileUploadReq, file)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the shardingResponse
fmt.Println("Object CID:", objectResponse.Data.ObjectCid)
fmt.Println("Object Name:", objectResponse.Data.ObjectName)
```

#### 4.3.15 ImportCarFileExt

Imports a Car file with an external reader.

```go
ImportCarFileExt(req *model.CarFileUploadReq, extReader io.Reader) (model.ObjectCreateResponse, error)
```

**Parameters**

| Name      | Type                    | Description                       |
|-----------|-------------------------|-----------------------------------|
| req       | *model.CarFileUploadReq | Car file upload request.          |
| extReader | io.Reader               | External reader for the Car file. |

**Return Value**

| Type                       | Description             |
|----------------------------|-------------------------|
| model.ObjectCreateResponse | Object create response. |
| error                      | Error, if any.          |

**Usage**

```go
carFileUploadReq := &model.CarFileUploadReq{
// Set the required fields in the Car file upload request
}

file, err := os.Open("path/to/car/file.car")
if err != nil {
fmt.Println("Error:", err)
return
}
defer file.Close()

objectResponse, err := sdk.Car.ImportCarFileExt(carFileUploadReq, file)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the objectResponse
fmt.Println("Object CID:", objectResponse.Data.ObjectCid)
fmt.Println("Object Name:", objectResponse.Data.ObjectName)
```

#### 4.3.16 ImportShardingCarFileExt

Imports a sharding Car file with an external reader.

```go
ImportShardingCarFileExt(req *model.CarFileUploadReq, extReader io.Reader) (model.ShardingCarFileUploadResponse, error)
```

**Parameters**

| Name      | Type                    | Description                       |
|-----------|-------------------------|-----------------------------------|
| req       | *model.CarFileUploadReq | Car file upload request.          |
| extReader | io.Reader               | External reader for the Car file. |

**Return Value**

| Type                                | Description                        |
|-------------------------------------|------------------------------------|
| model.ShardingCarFileUploadResponse | Sharding Car file upload response. |
| error                               | Error, if any.                     |

**Usage**

```go
carFileUploadReq := &model.CarFileUploadReq{
// Set the required fields in the Car file upload request
}

file, err := os.Open("path/to/car/file.car")
if err != nil {
fmt.Println("Error:", err)
return
}
defer file.Close()

shardingResponse, err := sdk.Car.ImportShardingCarFileExt(carFileUploadReq, file)
if err != nil {
fmt.Println("Error:", err)
return
}

// Process the shardingResponse
fmt.Println("Object CID:", objectResponse.Data.ObjectCid)
fmt.Println("Object Name:", objectResponse.Data.ObjectName)
```

#### 4.3.17 ExtractCarFile

Extracts a Car file.

```go
ExtractCarFile(carFilePath string, dataDestination string) error
```

**Parameters**

| Name            | Type   | Description                               |
|-----------------|--------|-------------------------------------------|
| carFilePath     | string | Path to the Car file.                     |
| dataDestination | string | Destination to extract the data from Car. |

**Return Value**

| Type  | Description    |
|-------|----------------|
| error | Error, if any. |

**Usage**

```go
carFilePath := "path/to/car/file.car"
dataDestination := "path/to/extract/data"

err := sdk.Car.ExtractCarFile(carFilePath, dataDestination)
if err != nil {
fmt.Println("Error:", err)
return
}

// Car file extracted successfully
```

### 4.4 Other

#### 4.4.1 GetIpfsVersion

Retrieves the version of IPFS (InterPlanetary File System).

```go
GetIpfsVersion() (model.VersionResponse, error)
```

**Return Value**

| Type                  | Description                              |
|-----------------------|------------------------------------------|
| model.VersionResponse | Version response containing the version. |
| error                 | Error, if any.                           |

**Usage**

```go
version, err := sdk.GetIpfsVersion()
if err != nil {
fmt.Println("Error:", err)
return
}

fmt.Println("IPFS Version:", version.Version)
```

#### 4.4.2 GetApiVersion

Retrieves the version of the Chainstorage-API.

```go
GetApiVersion() (model.VersionResponse, error)
```

**Return Value**

| Type                  | Description                              |
|-----------------------|------------------------------------------|
| model.VersionResponse | Version response containing the version. |
| error                 | Error, if any.                           |

**Usage**

```go
version, err := sdk.GetApiVersion()
if err != nil {
fmt.Println("Error:", err)
return
}

fmt.Println("Chainstorage API Version:", version.Version)
```

## 5. Troubleshooting

If you encounter any issues while using the SDK, you can refer to the following troubleshooting steps:

1. **Check SDK Compatibility:** Ensure that you are using a compatible version of the SDK with your application or
   system. Check the SDK documentation or release notes for compatibility information.

2. **Verify API Credentials:** Double-check the API credentials (e.g., API tokens) used for authentication. Make sure
   they are correct and have the necessary permissions to perform the desired operations.

3. **Check Network Connectivity:** Ensure that your application has a stable network connection. Check your network
   settings, firewall rules, and proxy configurations, if applicable.

4. **Review Error Messages:** When an error occurs, the SDK usually provides an error message or code. Review the error
   message to understand the issue better. You can consult the SDK documentation or API documentation for specific error
   code meanings and troubleshooting guidance.

5. **Validate Input Data:** Verify that the input data provided to the SDK methods is correct and in the expected
   format. Invalid or missing data can lead to errors or unexpected behavior.

6. **Handle Exceptions and Errors:** Properly handle exceptions and errors in your application code when interacting
   with the SDK. Implement error handling mechanisms such as try-catch blocks or error checks to gracefully handle
   exceptions and display meaningful error messages to the user.

7. **Update SDK Version:** If you are using an outdated version of the SDK, consider updating to the latest version.
   Newer versions often include bug fixes, performance improvements, and additional features.

8. **Consult SDK Documentation:** Refer to the SDK documentation for detailed usage instructions, code examples, and
   troubleshooting tips specific to the SDK.

9. **Contact Support:** If you are unable to resolve the issue using the above steps, reach out to the SDK support team
   or the API service provider for further assistance. Provide them with relevant information, such as error messages,
   request/response details, and steps to reproduce the issue.

By following these troubleshooting steps, you can diagnose and resolve common issues encountered while using the SDK and
ensure smooth integration with your application.

## 6. FAQ

**Q1: How do I authenticate with the API using the SDK?**
A: To authenticate with the API using the SDK, you typically need to provide your API credentials, such as
Chainstorage-API tokens, during the initialization or configuration of the SDK. Refer to the SDK documentation for
specific instructions on how to authenticate with the API using the SDK.

**Q2: Can I use multiple SDK instances in the same application?**
A: Yes, you can use multiple SDK instances in the same application if needed. Each SDK instance can be configured with
different settings or credentials to interact with different accounts, environments, or APIs. Ensure that you manage and
organize the SDK instances appropriately within your application.

**Q3: How can I handle rate limits and throttling with the SDK?**
A: The SDK typically provides built-in mechanisms to handle rate limits and throttling imposed by the API. It may
include automatic retries, exponential backoff, or error handling strategies. Consult the SDK documentation for details
on how the SDK handles rate limits and how you can customize or configure this behavior if necessary.

**Q4: Can I extend or customize the SDK functionality?**
A: Depending on the SDK design and capabilities, you may be able to extend or customize the SDK functionality. The SDK
may provide extension points, hooks, or customization options to modify its behavior or add additional functionality.
Refer to the SDK documentation for information on how to extend or customize the SDK.

**Q5: How can I contribute to the SDK development or report issues?**
A: If you would like to contribute to the SDK development or report issues, you can typically find the SDK's source code
repository or project page on platforms like [chainstorage-sdk](https://github.com/solarfs/go-chainstorage-sdk).
There, you can submit bug reports, feature requests, or even contribute code improvements through pull requests. Check
the SDK documentation or project page for specific instructions on how to contribute or report issues.

If you have any additional questions or need further assistance, please refer to the SDK documentation or contact the
SDK support team for more information.

## 7. Appendix

### 7.1 Glossary

#### Bucket

Represents a bucket.

**Fields**

| Field Name          | Type      | Description                                         |
|---------------------|-----------|-----------------------------------------------------|
| Id                  | int       | Bucket ID                                           |
| UserId              | int       | User ID                                             |
| BucketName          | string    | Bucket Name (3-63 character limit)                  |
| StorageNetworkCode  | int       | Storage Network Code (10001-IPFS)                   |
| BucketPrincipleCode | int       | Bucket Principle Code (10001-Public, 10000-Private) |
| UsedSpace           | int64     | Used Space (in bytes)                               |
| ObjectAmount        | int       | Object Amount                                       |
| Status              | int       | Record Status (0 - Active, 1 - Deleted)             |
| CreatedAt           | time.Time | Creation Time                                       |
| UpdatedAt           | time.Time | Last Update Time                                    |

#### BucketPageResponse

Represents a response containing a bucket page.

**Fields**

| Field Name | Type       | Description           |
|------------|------------|-----------------------|
| RequestId  | string     | Request ID (optional) |
| Code       | int32      | Code (optional)       |
| Msg        | string     | Message (optional)    |
| Status     | string     | Status (optional)     |
| Data       | BucketPage | Bucket Page Data      |

#### BucketPage

Represents a page of buckets.

**Fields**

| Field Name | Type     | Description                |
|------------|----------|----------------------------|
| Count      | int      | Count (optional)           |
| PageIndex  | int      | Page Index (optional)      |
| PageSize   | int      | Page Size (optional)       |
| List       | []Bucket | List of Buckets (optional) |

#### BucketCreateResponse

Represents a response containing a created bucket.

**Fields**

| Field Name | Type   | Description           |
|------------|--------|-----------------------|
| RequestId  | string | Request ID (optional) |
| Code       | int32  | Code (optional)       |
| Msg        | string | Message (optional)    |
| Status     | string | Status (optional)     |
| Data       | Bucket | Created Bucket Data   |

#### BucketEmptyResponse

Represents a response for emptying a bucket.

**Fields**

| Field Name | Type        | Description                |
|------------|-------------|----------------------------|
| RequestId  | string      | Request ID (optional)      |
| Code       | int32       | Code (optional)            |
| Msg        | string      | Message (optional)         |
| Status     | string      | Status (optional)          |
| Data       | interface{} | Additional Data (optional) |

#### BucketRemoveResponse

Represents a response for removing a bucket.

**Fields**

| Field Name | Type        | Description                |
|------------|-------------|----------------------------|
| RequestId  | string      | Request ID (optional)      |
| Code       | int32       | Code (optional)            |
| Msg        | string      | Message (optional)         |
| Status     | string      | Status (optional)          |
| Data       | interface{} | Additional Data (optional) |

#### Object

The `Object` struct represents an object and its attributes.

```go
type Object struct {
Id                      int                    `json:"id" comment:"对象ID"`
UserId                  int                    `json:"-" comment:"用户ID"`
BucketId                int                    `json:"bucketId" comment:"桶主键"`
ObjectName              string                 `json:"objectName" comment:"对象名称（255字限制）"`
ObjectTypeCode          int                    `json:"objectTypeCode" comment:"对象类型编码"`
ObjectSize              int64                  `json:"objectSize" comment:"对象大小（字节）"`
IsMarked                int                    `json:"isMarked" comment:"星标（1-已标记，0-未标记）"`
ObjectCid               string                 `json:"objectCid" comment:"对象CID"`
Status                  int                    `json:"status" comment:"记录状态（0-有效，1-删除）"`
LinkedStorageObjectCode string                 `json:"linkedStorageObjectCode" comment:"链存类型对象编码"`
LinkedStorageObject     map[string]interface{} `json:"linkedStorageObject" comment:"链存类型对象"`
CreatedAt               time.Time              `json:"createdAt" comment:"创建时间"`
UpdatedAt               time.Time              `json:"updatedAt" comment:"最后更新时间"`
}
```

#### ObjectPageResponse

The `ObjectPageResponse` struct represents a response containing a page of objects.

```go
type ObjectPageResponse struct {
RequestId string     `json:"requestId,omitempty"`
Code      int32      `json:"code,omitempty"`
Msg       string     `json:"msg,omitempty"`
Status    string     `json:"status,omitempty"`
Data      ObjectPage `json:"data,omitempty"`
}
```

#### ObjectPage

The `ObjectPage` struct represents a page of objects.

```go
type ObjectPage struct {
Count     int      `json:"count,omitempty"`
PageIndex int      `json:"pageIndex,omitempty"`
PageSize  int      `json:"pageSize,omitempty"`
List      []Object `json:"list,omitempty"`
}
```

#### ObjectCreateResponse

The `ObjectCreateResponse` struct represents a response after creating an object.

```go
type ObjectCreateResponse struct {
RequestId string `json:"requestId,omitempty"`
Code      int32  `json:"code,omitempty"`
Msg       string `json:"msg,omitempty"`
Status    string `json:"status,omitempty"`
Data      Object `json:"data,omitempty"`
}
```

#### ObjectRemoveResponse

The `ObjectRemoveResponse` struct represents a response after removing an object.

```go
type ObjectRemoveResponse struct {
RequestId string      `json:"requestId,omitempty"`
Code      int32       `json:"code,omitempty"`
Msg       string      `json:"msg,omitempty"`
Status    string      `json:"status,omitempty"`
Data      interface{} `json:"data,omitempty"`
}
```

#### ObjectRenameResponse

The `ObjectRenameResponse` struct represents a response after renaming an object.

```go
type ObjectRenameResponse struct {
RequestId string      `json:"requestId,omitempty"`
Code      int32       `json:"code,omitempty"`
Msg       string      `json:"msg,omitempty"`
Status    string      `json:"status,omitempty"`
Data      interface{} `json:"data,omitempty"`
}
```

#### ObjectMarkResponse

The `ObjectMarkResponse` struct represents a response after marking an object.

```go
type ObjectMarkResponse struct {
RequestId string      `json:"requestId,omitempty"`
Code      int32       `json:"code,omitempty"`
Msg       string      `json:"msg,omitempty"`
Status    string      `json:"status,omitempty"`


Data      interface{} `json:"data,omitempty"`
}
```

#### ObjectExistResponse

The `ObjectExistResponse` struct represents a response after checking the existence of an object.

```go
type ObjectExistResponse struct {
RequestId string           `json:"requestId,omitempty"`
Code      int32            `json:"code,omitempty"`
Msg       string           `json:"msg,omitempty"`
Status    string           `json:"status,omitempty"`
Data      ObjectExistCheck `json:"data,omitempty"`
}
```

#### ObjectExistCheck

The `ObjectExistCheck` struct represents the result of object existence check.

```go
type ObjectExistCheck struct {
IsExist bool `json:"isExist"`
}
```

#### CarFileUploadReq

Represents the request for uploading a car file.

**Fields**

| Field           | Type   | Description                            |
|-----------------|--------|----------------------------------------|
| BucketId        | int    | The bucket ID (required)               |
| ObjectName      | string | The object name (255 characters limit) |
| ObjectTypeCode  | int    | The object type code                   |
| ObjectSize      | int64  | The object size in bytes               |
| ObjectCid       | string | The object CID                         |
| FileDestination | string | The file path                          |
| RawSha256       | string | The raw file SHA256 hash (required)    |
| ShardingSha256  | string | The sharding file SHA256 hash          |
| ShardingNo      | int    | The sharding sequence number           |
| ShardingAmount  | int    | The total number of sharding files     |
| CarFileCid      | string | The car file CID                       |

#### ShardingCarFileUploadResponse

Represents the response for sharding car file upload.

**Fields**

| Field     | Type          | Description                       |
|-----------|---------------|-----------------------------------|
| RequestId | string        | The request ID                    |
| Code      | int32         | The response code                 |
| Msg       | string        | The response message              |
| Status    | string        | The response status               |
| Data      | CarFileUpload | The uploaded car file information |

#### CarFileUpload

Represents the uploaded car file information.

**Fields**

| Field          | Type   | Description                            |
|----------------|--------|----------------------------------------|
| BucketId       | int    | The bucket ID (required)               |
| ObjectName     | string | The object name (255 characters limit) |
| ObjectTypeCode | int    | The object type code                   |
| ObjectSize     | int64  | The object size in bytes               |
| ObjectCid      | string | The object CID                         |
| RawSha256      | string | The raw file SHA256 hash (required)    |
| ShardingSha256 | string | The sharding file SHA256 hash          |
| ShardingNo     | int    | The sharding sequence number           |
| ShardingAmount | int    | The total number of sharding files     |
| CarFileCid     | string | The car file CID                       |

#### ShardingCarFilesVerifyResponse

Represents the response for verifying sharding car files.

**Fields**

| Field          | Type     | Description                         |
|----------------|----------|-------------------------------------|
| RawSha256      | string   | The raw file SHA256 hash (required) |
| ObjectName     | string   | The file name                       |
| ShardingAmount | int      | The total number of sharding files  |
| UploadStatus   | int      | The upload status                   |
| Uploaded       | []string | The list of uploaded sharding files |

#### CarResponse

Represents a general car response.

**Fields**

| Field     | Type        | Description          |
|-----------|-------------|----------------------|
| RequestId | string      | The request ID       |
| Code      | int32       | The response code    |
| Msg       | string      | The response message |
| Status    | string      | The response status  |
| Data      | interface{} | The response data    |

#### RootLink

Represents the root link.

**Fields**

| Field       | Type    | Description                  |
|-------------|---------|------------------------------|
| (Inherited) |         | Inherits from `ipldfmt.Link` |
| RootCid     | cid.Cid | The root CID                 |

### 7.2 Error Codes

| Error Code | Description                                                                                                                | 描述                                     |
|------------|----------------------------------------------------------------------------------------------------------------------------|----------------------------------------|
| 100201     | This bucket data does not exist                                                                                            | 该数据桶不存在                                |
| 100202     | The object data is not in the bucket                                                                                       | 桶内没有该对象数据                              |
| 100203     | Bucket name must be between 3-63 characters and can only contain lowercase letters, numbers, and hyphens, please try again | 桶名称必须是3-63个字符之间的小写字母、数字和破折号，请重新尝试      |
| 100204     | Bucket name is already taken, please choose a different name                                                               | 桶名称已被占用，请更换桶名称                         |
| 100205     | Incorrect storage network name settings, please try again                                                                  | 存储网络名称设置不正确，请重新尝试                      |
| 100206     | Please set the bucket policy correctly                                                                                     | 请正确设置桶策略                               |
| 100207     | The bucket contains data and cannot be deleted                                                                             | 桶内有数据，无法删除                             |
| 100208     | Error occurred while calculating bucket capacity, please try again                                                         | 桶容量统计出错，请重试                            |
| 100209     | Error occurred while getting bucket quota, please try again                                                                | 桶容量配额获取出错，请重试                          |
| 100210     | Error occurred while updating bucket quota, please try again                                                               | 桶容量配额更新出错，请重试                          |
| 100211     | In the basic version, only one bucket can be created for each network type                                                 | 基础版本限制，每种网络类型只能创建一个桶                   |
| 100212     | Invalid bucket ID                                                                                                          | 桶ID无效                                  |
| 100213     | Incorrect storage network code settings, please try again                                                                  | 存储网络编码设置不正确，请重新尝试                      |
| 100214     | The bucket is bound with a gateway and cannot be deleted                                                                   | 桶已经与网关绑定，无法删除                          |
| 100301     | The object data does not exist                                                                                             | 该对象数据不存在                               |
| 100302     | Object name must be between 1-255 characters and cannot contain invalid characters, please try again                       | 对象名称必须是1-255个字符之间，不能包含非法字符，请重试         |
| 100303     | Object name already exists, do you want to overwrite the existing object?                                                  | 对象名称已存在，是否覆盖原有对象                       |
| 100304     | Error occurred while operating the object reference counter                                                                | 对象引用计数器操作出错                            |
| 100305     | Invalid object CID                                                                                                         | 无效的对象CID                               |
| 100306     | Invalid object ID                                                                                                          | 无效的对象ID                                |
| 100307     | Invalid object ID list                                                                                                     | 无效的对象ID列表                              |
| 100801     | The APIKey does not exist                                                                                                  | 该APIKey不存在                             |
| 100802     | APIKey name must be between 3-63 characters and can only contain lowercase letters, numbers, and hyphens, please try again | APIKey名称必须是3-63个字符之间的小写字母、数字和破折号，请重新尝试 |
| 100803     | APIKey name already exists, please try again                                                                               | APIKey名称已存在，请重新尝试                      |
| 100804     | Failed to create APIKey, please try again                                                                                  | APIKey创建失败，请重试                         |
| 100805     | Incorrect admin settings, please try again                                                                                 | 管理员设置不正确，请重试                           |
| 100806     | APIKey permissions settings are incorrect, please try again                                                                | APIKey权限设置错误，请重试                       |
| 100807     | Incorrect APIKey data range setting, please try again                                                                      | APIKey数据范围设置不正确，请重试                    |
| 100808     | Incorrect PinningServiceAPI permissions setting, please try again                                                          | PinningServiceAPI权限设置不正确，请重试           |
| 100901     | Fail to upload CAR file                                                                                                    | CAR上传文件失败                              |
| 100902     | Invalid uploading data path                                                                                                | 无效的上传数据路径                              |
| 100903     | Fail to create CAR file                                                                                                    | 创建CAR文件失败                              |
| 100904     | Fail to parse CAR file                                                                                                     | 解析CAR文件失败                              |
| 100905     | Fail to compute CAR file HASH                                                                                              | CAR文件HASH计算失败                          |
| 100906     | Fail to chunk CAR file                                                                                                     | 生成CAR文件分片操作失败                          |
| 100907     | Fail to reference object by CID                                                                                            | 执行CID秒传操作失败                            |
| 100908     | CID of the CAR file does not match the raw CID                                                                             | CAR文件的CID与原始CID不匹配                     |
| 100909     | Uploading folder is empty, or uploading data is invalid in the folder                                                      | 上传目录为空或者目录中的数据无效                       |
| 100910     | Exceed the limitation of object amount                                                                                     | 超过对象存储限制                               |
| 100911     | Exceed the limitation of storage space                                                                                     | 超过空间存储限制                               |
| 100912     | Exceed the limitation of items in the upload folder                                                                        | 超过上传文件夹的项目限制                           |


