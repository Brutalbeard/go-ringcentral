# PhoneNumberInfoNumberParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaCode** | **string** | Area code of location. The portion of the [E.164](https://www.itu.int/rec/T-REC-E.164-201011-I) number that identifies a specific geographic region/numbering area of the national numbering plan (NANP); that can be summarized as &#x60;NPA-NXX-xxxx&#x60; and covers Canada, the United States, parts of the Caribbean Sea, and some Atlantic and Pacific islands. See [North American Numbering Plan] (https://en.wikipedia.org/wiki/North_American_Numbering_Plan) for details | [optional] 
**Country** | [**[]GetCountryInfoNumberParser**](GetCountryInfoNumberParser.md) | Information on a country the phone number belongs to | [optional] 
**Dialable** | **string** | Dialing format of a phone number | [optional] 
**E164** | **string** | Phone number [E.164](https://www.itu.int/rec/T-REC-E.164-201011-I) format | [optional] 
**FormattedInternational** | **string** | International format of a phone number | [optional] 
**FormattedNational** | **string** | Domestic format of a phone number | [optional] 
**OriginalString** | **string** | One of the numbers to be parsed, passed as a string in response | [optional] 
**Special** | **bool** | &#39;True&#39; if the number is in a special format (for example N11 code) | [optional] 
**Normalized** | **string** | Phone number [E.164](https://www.itu.int/rec/T-REC-E.164-201011-I) format without plus sign (&#39;+&#39;) | [optional] 
**TollFree** | **bool** | Specifies if a phone number is toll free or not | [optional] 
**SubAddress** | **string** | Sub-Address. The portion of the number that identifies a subscriber into the subscriber internal (non-public) network. | [optional] 
**SubscriberNumber** | **string** |  Subscriber number. The portion of the [E.164](https://www.itu.int/rec/T-REC-E.164-201011-I) number that identifies a subscriber in a network or numbering area. | [optional] 
**DtmfPostfix** | **string** | DTMF (Dual Tone Multi-Frequency) postfix | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


