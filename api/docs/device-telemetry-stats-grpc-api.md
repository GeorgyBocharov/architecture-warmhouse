# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [telemetry.proto](#telemetry-proto)
    - [DataPoint](#warmhouse-telemetry-v1-DataPoint)
    - [Device](#warmhouse-telemetry-v1-Device)
    - [GetTelemetryRequest](#warmhouse-telemetry-v1-GetTelemetryRequest)
    - [House](#warmhouse-telemetry-v1-House)
    - [SensorData](#warmhouse-telemetry-v1-SensorData)
    - [TelemetryResponse](#warmhouse-telemetry-v1-TelemetryResponse)
    - [TimeRange](#warmhouse-telemetry-v1-TimeRange)
  
    - [TelemetryService](#warmhouse-telemetry-v1-TelemetryService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="telemetry-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## telemetry.proto



<a name="warmhouse-telemetry-v1-DataPoint"></a>

### DataPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | required |
| value | [int64](#int64) |  | required |






<a name="warmhouse-telemetry-v1-Device"></a>

### Device



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | required |
| name | [string](#string) |  | required |
| room_name | [string](#string) |  | optional |
| house | [House](#warmhouse-telemetry-v1-House) |  | optional |






<a name="warmhouse-telemetry-v1-GetTelemetryRequest"></a>

### GetTelemetryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device_id | [string](#string) |  | required |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | required |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | optional |
| sensor_ids | [string](#string) | repeated |  |
| user_id | [string](#string) |  | required |






<a name="warmhouse-telemetry-v1-House"></a>

### House



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | required |
| name | [string](#string) |  | optional |
| address | [string](#string) |  | optional |






<a name="warmhouse-telemetry-v1-SensorData"></a>

### SensorData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | required |
| type | [string](#string) |  | required |
| unit | [string](#string) |  | required |
| data_points | [DataPoint](#warmhouse-telemetry-v1-DataPoint) | repeated |  |






<a name="warmhouse-telemetry-v1-TelemetryResponse"></a>

### TelemetryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [Device](#warmhouse-telemetry-v1-Device) |  | required |
| time_range | [TimeRange](#warmhouse-telemetry-v1-TimeRange) |  | required |
| sensor_data | [SensorData](#warmhouse-telemetry-v1-SensorData) | repeated |  |






<a name="warmhouse-telemetry-v1-TimeRange"></a>

### TimeRange



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | required |
| end | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | required |





 

 

 


<a name="warmhouse-telemetry-v1-TelemetryService"></a>

### TelemetryService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetTelemetry | [GetTelemetryRequest](#warmhouse-telemetry-v1-GetTelemetryRequest) | [TelemetryResponse](#warmhouse-telemetry-v1-TelemetryResponse) | Получить статистику телеметрии устройства |

 



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

