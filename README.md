# Go API client for libregraph

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v0.13.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import libregraph "github.com/owncloud/libre-graph-api-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), libregraph.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), libregraph.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), libregraph.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), libregraph.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DrivesApi* | [**CreateDrive**](docs/DrivesApi.md#createdrive) | **Post** /drives | Create a new space of a specific type
*DrivesApi* | [**DeleteDrive**](docs/DrivesApi.md#deletedrive) | **Delete** /drives/{drive-id} | Delete a specific space
*DrivesApi* | [**GetDrive**](docs/DrivesApi.md#getdrive) | **Get** /drives/{drive-id} | Get drive by id
*DrivesApi* | [**UpdateDrive**](docs/DrivesApi.md#updatedrive) | **Patch** /drives/{drive-id} | Update the space
*DrivesGetDrivesApi* | [**ListAllDrives**](docs/DrivesGetDrivesApi.md#listalldrives) | **Get** /drives | Get All drives
*DrivesRootApi* | [**GetRoot**](docs/DrivesRootApi.md#getroot) | **Get** /drives/{drive-id}/root | Get root from arbitrary space
*GroupApi* | [**AddMember**](docs/GroupApi.md#addmember) | **Post** /groups/{group-id}/members/$ref | Add a member to a group
*GroupApi* | [**DeleteGroup**](docs/GroupApi.md#deletegroup) | **Delete** /groups/{group-id} | Delete entity from groups
*GroupApi* | [**DeleteMember**](docs/GroupApi.md#deletemember) | **Delete** /groups/{group-id}/members/{directory-object-id}/$ref | Delete member from a group
*GroupApi* | [**GetGroup**](docs/GroupApi.md#getgroup) | **Get** /groups/{group-id} | Get entity from groups by key
*GroupApi* | [**UpdateGroup**](docs/GroupApi.md#updategroup) | **Patch** /groups/{group-id} | Update entity in groups
*GroupsApi* | [**CreateGroup**](docs/GroupsApi.md#creategroup) | **Post** /groups | Add new entity to groups
*GroupsApi* | [**ListGroups**](docs/GroupsApi.md#listgroups) | **Get** /groups | Get entities from groups
*MeDriveApi* | [**GetHome**](docs/MeDriveApi.md#gethome) | **Get** /me/drive | Get personal space for user
*MeDriveRootApi* | [**HomeGetRoot**](docs/MeDriveRootApi.md#homegetroot) | **Get** /me/drive/root | Get root from personal space
*MeDriveRootChildrenApi* | [**HomeGetChildren**](docs/MeDriveRootChildrenApi.md#homegetchildren) | **Get** /me/drive/root/children | Get children from drive
*MeDrivesApi* | [**ListMyDrives**](docs/MeDrivesApi.md#listmydrives) | **Get** /me/drives | Get drives from me
*MeUserApi* | [**MeGet**](docs/MeUserApi.md#meget) | **Get** /me | 
*UserApi* | [**DeleteUser**](docs/UserApi.md#deleteuser) | **Delete** /users/{user-id} | Delete entity from users
*UserApi* | [**GetUser**](docs/UserApi.md#getuser) | **Get** /users/{user-id} | Get entity from users by key
*UserApi* | [**UpdateUser**](docs/UserApi.md#updateuser) | **Patch** /users/{user-id} | Update entity in users
*UsersApi* | [**CreateUser**](docs/UsersApi.md#createuser) | **Post** /users | Add new entity to users
*UsersApi* | [**ListUsers**](docs/UsersApi.md#listusers) | **Get** /users | Get entities from users


## Documentation For Models

 - [CollectionOfDriveItems](docs/CollectionOfDriveItems.md)
 - [CollectionOfDrives](docs/CollectionOfDrives.md)
 - [CollectionOfGroup](docs/CollectionOfGroup.md)
 - [CollectionOfUser](docs/CollectionOfUser.md)
 - [Deleted](docs/Deleted.md)
 - [DirectoryObject](docs/DirectoryObject.md)
 - [Drive](docs/Drive.md)
 - [DriveItem](docs/DriveItem.md)
 - [FileSystemInfo](docs/FileSystemInfo.md)
 - [Folder](docs/Folder.md)
 - [FolderView](docs/FolderView.md)
 - [Group](docs/Group.md)
 - [Hashes](docs/Hashes.md)
 - [Identity](docs/Identity.md)
 - [IdentitySet](docs/IdentitySet.md)
 - [Image](docs/Image.md)
 - [ItemReference](docs/ItemReference.md)
 - [MemberReference](docs/MemberReference.md)
 - [OdataError](docs/OdataError.md)
 - [OdataErrorDetail](docs/OdataErrorDetail.md)
 - [OdataErrorMain](docs/OdataErrorMain.md)
 - [OpenGraphFile](docs/OpenGraphFile.md)
 - [PasswordProfile](docs/PasswordProfile.md)
 - [Permission](docs/Permission.md)
 - [Quota](docs/Quota.md)
 - [RemoteItem](docs/RemoteItem.md)
 - [Shared](docs/Shared.md)
 - [SpecialFolder](docs/SpecialFolder.md)
 - [Trash](docs/Trash.md)
 - [User](docs/User.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



