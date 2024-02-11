# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [game/v1/game.proto](#game_v1_game-proto)
    - [EndSessionRequest](#game-v1-EndSessionRequest)
    - [EndSessionResponse](#game-v1-EndSessionResponse)
    - [PlayGameRequest](#game-v1-PlayGameRequest)
    - [PlayGameResponse](#game-v1-PlayGameResponse)
    - [Session](#game-v1-Session)
    - [StartSessionRequest](#game-v1-StartSessionRequest)
    - [StartSessionResponse](#game-v1-StartSessionResponse)
  
    - [Move](#game-v1-Move)
  
- [game/v1/service.proto](#game_v1_service-proto)
    - [GameService](#game-v1-GameService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="game_v1_game-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## game/v1/game.proto



<a name="game-v1-EndSessionRequest"></a>

### EndSessionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session | [Session](#game-v1-Session) |  |  |






<a name="game-v1-EndSessionResponse"></a>

### EndSessionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session | [Session](#game-v1-Session) |  |  |






<a name="game-v1-PlayGameRequest"></a>

### PlayGameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [string](#string) |  |  |
| player_move | [Move](#game-v1-Move) |  |  |






<a name="game-v1-PlayGameResponse"></a>

### PlayGameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| computer_move | [Move](#game-v1-Move) |  |  |
| session | [Session](#game-v1-Session) |  |  |






<a name="game-v1-Session"></a>

### Session



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| player_id | [string](#string) |  |  |
| wins | [int32](#int32) |  |  |
| losses | [int32](#int32) |  |  |
| draws | [int32](#int32) |  |  |






<a name="game-v1-StartSessionRequest"></a>

### StartSessionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| player_id | [string](#string) | optional |  |






<a name="game-v1-StartSessionResponse"></a>

### StartSessionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session | [Session](#game-v1-Session) |  |  |





 


<a name="game-v1-Move"></a>

### Move


| Name | Number | Description |
| ---- | ------ | ----------- |
| MOVE_UNSPECIFIED | 0 |  |
| MOVE_ROCK | 1 |  |
| MOVE_PAPER | 2 |  |
| MOVE_SCISSORS | 3 |  |


 

 

 



<a name="game_v1_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## game/v1/service.proto


 

 

 


<a name="game-v1-GameService"></a>

### GameService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| StartSession | [StartSessionRequest](#game-v1-StartSessionRequest) | [StartSessionResponse](#game-v1-StartSessionResponse) |  |
| PlayGame | [PlayGameRequest](#game-v1-PlayGameRequest) | [PlayGameResponse](#game-v1-PlayGameResponse) |  |
| EndSession | [EndSessionRequest](#game-v1-EndSessionRequest) | [EndSessionResponse](#game-v1-EndSessionResponse) |  |

 



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

