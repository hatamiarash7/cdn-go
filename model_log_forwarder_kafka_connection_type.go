/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.103.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the LogForwarderKafkaConnectionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogForwarderKafkaConnectionType{}

// LogForwarderKafkaConnectionType Kafka connection
type LogForwarderKafkaConnectionType struct {
	SampleRate *int32 `json:"sample_rate,omitempty"`
	KafkaVersion *string `json:"kafka_version,omitempty"`
	KafkaBrokers []string `json:"kafka_brokers,omitempty"`
	KafkaTopicToWrite *string `json:"kafka_topic_to_write,omitempty"`
	KafkaProducerBatchSize *int32 `json:"kafka_producer_batch_size,omitempty"`
	KafkaProducerFlushFrequencyMs *int32 `json:"kafka_producer_flush_frequency_ms,omitempty"`
}

// NewLogForwarderKafkaConnectionType instantiates a new LogForwarderKafkaConnectionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwarderKafkaConnectionType() *LogForwarderKafkaConnectionType {
	this := LogForwarderKafkaConnectionType{}
	return &this
}

// NewLogForwarderKafkaConnectionTypeWithDefaults instantiates a new LogForwarderKafkaConnectionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwarderKafkaConnectionTypeWithDefaults() *LogForwarderKafkaConnectionType {
	this := LogForwarderKafkaConnectionType{}
	return &this
}

// GetSampleRate returns the SampleRate field value if set, zero value otherwise.
func (o *LogForwarderKafkaConnectionType) GetSampleRate() int32 {
	if o == nil || IsNil(o.SampleRate) {
		var ret int32
		return ret
	}
	return *o.SampleRate
}

// GetSampleRateOk returns a tuple with the SampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderKafkaConnectionType) GetSampleRateOk() (*int32, bool) {
	if o == nil || IsNil(o.SampleRate) {
		return nil, false
	}
	return o.SampleRate, true
}

// HasSampleRate returns a boolean if a field has been set.
func (o *LogForwarderKafkaConnectionType) HasSampleRate() bool {
	if o != nil && !IsNil(o.SampleRate) {
		return true
	}

	return false
}

// SetSampleRate gets a reference to the given int32 and assigns it to the SampleRate field.
func (o *LogForwarderKafkaConnectionType) SetSampleRate(v int32) {
	o.SampleRate = &v
}

// GetKafkaVersion returns the KafkaVersion field value if set, zero value otherwise.
func (o *LogForwarderKafkaConnectionType) GetKafkaVersion() string {
	if o == nil || IsNil(o.KafkaVersion) {
		var ret string
		return ret
	}
	return *o.KafkaVersion
}

// GetKafkaVersionOk returns a tuple with the KafkaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderKafkaConnectionType) GetKafkaVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KafkaVersion) {
		return nil, false
	}
	return o.KafkaVersion, true
}

// HasKafkaVersion returns a boolean if a field has been set.
func (o *LogForwarderKafkaConnectionType) HasKafkaVersion() bool {
	if o != nil && !IsNil(o.KafkaVersion) {
		return true
	}

	return false
}

// SetKafkaVersion gets a reference to the given string and assigns it to the KafkaVersion field.
func (o *LogForwarderKafkaConnectionType) SetKafkaVersion(v string) {
	o.KafkaVersion = &v
}

// GetKafkaBrokers returns the KafkaBrokers field value if set, zero value otherwise.
func (o *LogForwarderKafkaConnectionType) GetKafkaBrokers() []string {
	if o == nil || IsNil(o.KafkaBrokers) {
		var ret []string
		return ret
	}
	return o.KafkaBrokers
}

// GetKafkaBrokersOk returns a tuple with the KafkaBrokers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderKafkaConnectionType) GetKafkaBrokersOk() ([]string, bool) {
	if o == nil || IsNil(o.KafkaBrokers) {
		return nil, false
	}
	return o.KafkaBrokers, true
}

// HasKafkaBrokers returns a boolean if a field has been set.
func (o *LogForwarderKafkaConnectionType) HasKafkaBrokers() bool {
	if o != nil && !IsNil(o.KafkaBrokers) {
		return true
	}

	return false
}

// SetKafkaBrokers gets a reference to the given []string and assigns it to the KafkaBrokers field.
func (o *LogForwarderKafkaConnectionType) SetKafkaBrokers(v []string) {
	o.KafkaBrokers = v
}

// GetKafkaTopicToWrite returns the KafkaTopicToWrite field value if set, zero value otherwise.
func (o *LogForwarderKafkaConnectionType) GetKafkaTopicToWrite() string {
	if o == nil || IsNil(o.KafkaTopicToWrite) {
		var ret string
		return ret
	}
	return *o.KafkaTopicToWrite
}

// GetKafkaTopicToWriteOk returns a tuple with the KafkaTopicToWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderKafkaConnectionType) GetKafkaTopicToWriteOk() (*string, bool) {
	if o == nil || IsNil(o.KafkaTopicToWrite) {
		return nil, false
	}
	return o.KafkaTopicToWrite, true
}

// HasKafkaTopicToWrite returns a boolean if a field has been set.
func (o *LogForwarderKafkaConnectionType) HasKafkaTopicToWrite() bool {
	if o != nil && !IsNil(o.KafkaTopicToWrite) {
		return true
	}

	return false
}

// SetKafkaTopicToWrite gets a reference to the given string and assigns it to the KafkaTopicToWrite field.
func (o *LogForwarderKafkaConnectionType) SetKafkaTopicToWrite(v string) {
	o.KafkaTopicToWrite = &v
}

// GetKafkaProducerBatchSize returns the KafkaProducerBatchSize field value if set, zero value otherwise.
func (o *LogForwarderKafkaConnectionType) GetKafkaProducerBatchSize() int32 {
	if o == nil || IsNil(o.KafkaProducerBatchSize) {
		var ret int32
		return ret
	}
	return *o.KafkaProducerBatchSize
}

// GetKafkaProducerBatchSizeOk returns a tuple with the KafkaProducerBatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderKafkaConnectionType) GetKafkaProducerBatchSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.KafkaProducerBatchSize) {
		return nil, false
	}
	return o.KafkaProducerBatchSize, true
}

// HasKafkaProducerBatchSize returns a boolean if a field has been set.
func (o *LogForwarderKafkaConnectionType) HasKafkaProducerBatchSize() bool {
	if o != nil && !IsNil(o.KafkaProducerBatchSize) {
		return true
	}

	return false
}

// SetKafkaProducerBatchSize gets a reference to the given int32 and assigns it to the KafkaProducerBatchSize field.
func (o *LogForwarderKafkaConnectionType) SetKafkaProducerBatchSize(v int32) {
	o.KafkaProducerBatchSize = &v
}

// GetKafkaProducerFlushFrequencyMs returns the KafkaProducerFlushFrequencyMs field value if set, zero value otherwise.
func (o *LogForwarderKafkaConnectionType) GetKafkaProducerFlushFrequencyMs() int32 {
	if o == nil || IsNil(o.KafkaProducerFlushFrequencyMs) {
		var ret int32
		return ret
	}
	return *o.KafkaProducerFlushFrequencyMs
}

// GetKafkaProducerFlushFrequencyMsOk returns a tuple with the KafkaProducerFlushFrequencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderKafkaConnectionType) GetKafkaProducerFlushFrequencyMsOk() (*int32, bool) {
	if o == nil || IsNil(o.KafkaProducerFlushFrequencyMs) {
		return nil, false
	}
	return o.KafkaProducerFlushFrequencyMs, true
}

// HasKafkaProducerFlushFrequencyMs returns a boolean if a field has been set.
func (o *LogForwarderKafkaConnectionType) HasKafkaProducerFlushFrequencyMs() bool {
	if o != nil && !IsNil(o.KafkaProducerFlushFrequencyMs) {
		return true
	}

	return false
}

// SetKafkaProducerFlushFrequencyMs gets a reference to the given int32 and assigns it to the KafkaProducerFlushFrequencyMs field.
func (o *LogForwarderKafkaConnectionType) SetKafkaProducerFlushFrequencyMs(v int32) {
	o.KafkaProducerFlushFrequencyMs = &v
}

func (o LogForwarderKafkaConnectionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogForwarderKafkaConnectionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SampleRate) {
		toSerialize["sample_rate"] = o.SampleRate
	}
	if !IsNil(o.KafkaVersion) {
		toSerialize["kafka_version"] = o.KafkaVersion
	}
	if !IsNil(o.KafkaBrokers) {
		toSerialize["kafka_brokers"] = o.KafkaBrokers
	}
	if !IsNil(o.KafkaTopicToWrite) {
		toSerialize["kafka_topic_to_write"] = o.KafkaTopicToWrite
	}
	if !IsNil(o.KafkaProducerBatchSize) {
		toSerialize["kafka_producer_batch_size"] = o.KafkaProducerBatchSize
	}
	if !IsNil(o.KafkaProducerFlushFrequencyMs) {
		toSerialize["kafka_producer_flush_frequency_ms"] = o.KafkaProducerFlushFrequencyMs
	}
	return toSerialize, nil
}

type NullableLogForwarderKafkaConnectionType struct {
	value *LogForwarderKafkaConnectionType
	isSet bool
}

func (v NullableLogForwarderKafkaConnectionType) Get() *LogForwarderKafkaConnectionType {
	return v.value
}

func (v *NullableLogForwarderKafkaConnectionType) Set(val *LogForwarderKafkaConnectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderKafkaConnectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderKafkaConnectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderKafkaConnectionType(val *LogForwarderKafkaConnectionType) *NullableLogForwarderKafkaConnectionType {
	return &NullableLogForwarderKafkaConnectionType{value: val, isSet: true}
}

func (v NullableLogForwarderKafkaConnectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderKafkaConnectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


