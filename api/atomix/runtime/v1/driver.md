# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/runtime/v1/driver.proto](#atomix_runtime_v1_driver-proto)
    - [Driver](#atomix-management-v1-Driver)
    - [DriverChunk](#atomix-management-v1-DriverChunk)
    - [DriverHeader](#atomix-management-v1-DriverHeader)
    - [DriverId](#atomix-management-v1-DriverId)
    - [DriverMeta](#atomix-management-v1-DriverMeta)
    - [DriverTrailer](#atomix-management-v1-DriverTrailer)
    - [InstallDriverRequest](#atomix-management-v1-InstallDriverRequest)
    - [InstallDriverResponse](#atomix-management-v1-InstallDriverResponse)
  
    - [DriverService](#atomix-management-v1-DriverService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_runtime_v1_driver-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/runtime/v1/driver.proto



<a name="atomix-management-v1-Driver"></a>

### Driver



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [DriverMeta](#atomix-management-v1-DriverMeta) |  |  |






<a name="atomix-management-v1-DriverChunk"></a>

### DriverChunk



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  |  |






<a name="atomix-management-v1-DriverHeader"></a>

### DriverHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| driver | [Driver](#atomix-management-v1-Driver) |  |  |






<a name="atomix-management-v1-DriverId"></a>

### DriverId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| version | [string](#string) |  |  |






<a name="atomix-management-v1-DriverMeta"></a>

### DriverMeta



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [DriverId](#atomix-management-v1-DriverId) |  |  |
| version | [uint64](#uint64) |  |  |






<a name="atomix-management-v1-DriverTrailer"></a>

### DriverTrailer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| checksum | [string](#string) |  |  |






<a name="atomix-management-v1-InstallDriverRequest"></a>

### InstallDriverRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [DriverHeader](#atomix-management-v1-DriverHeader) |  |  |
| chunk | [DriverChunk](#atomix-management-v1-DriverChunk) |  |  |
| trailer | [DriverTrailer](#atomix-management-v1-DriverTrailer) |  |  |






<a name="atomix-management-v1-InstallDriverResponse"></a>

### InstallDriverResponse






 

 

 


<a name="atomix-management-v1-DriverService"></a>

### DriverService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| InstallDriver | [InstallDriverRequest](#atomix-management-v1-InstallDriverRequest) stream | [InstallDriverResponse](#atomix-management-v1-InstallDriverResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
