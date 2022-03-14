/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.18
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// FiltersRouteTable One or more filters.
type FiltersRouteTable struct {
	// The IDs of the route tables involved in the associations.
	LinkRouteTableIds *[]string `json:"LinkRouteTableIds,omitempty"`
	// The IDs of the associations between the route tables and the Subnets.
	LinkRouteTableLinkRouteTableIds *[]string `json:"LinkRouteTableLinkRouteTableIds,omitempty"`
	// If true, the route tables are the main ones for their Nets.
	LinkRouteTableMain *bool `json:"LinkRouteTableMain,omitempty"`
	// The IDs of the Subnets involved in the associations.
	LinkSubnetIds *[]string `json:"LinkSubnetIds,omitempty"`
	// The IDs of the Nets for the route tables.
	NetIds *[]string `json:"NetIds,omitempty"`
	// The methods used to create a route.
	RouteCreationMethods *[]string `json:"RouteCreationMethods,omitempty"`
	// The IP ranges specified in routes in the tables.
	RouteDestinationIpRanges *[]string `json:"RouteDestinationIpRanges,omitempty"`
	// The service IDs specified in routes in the tables.
	RouteDestinationServiceIds *[]string `json:"RouteDestinationServiceIds,omitempty"`
	// The IDs of the gateways specified in routes in the tables.
	RouteGatewayIds *[]string `json:"RouteGatewayIds,omitempty"`
	// The IDs of the NAT services specified in routes in the tables.
	RouteNatServiceIds *[]string `json:"RouteNatServiceIds,omitempty"`
	// The IDs of the Net peering connections specified in routes in the tables.
	RouteNetPeeringIds *[]string `json:"RouteNetPeeringIds,omitempty"`
	// The states of routes in the route tables (`active` \\| `blackhole`). The `blackhole` state indicates that the target of the route is not available.
	RouteStates *[]string `json:"RouteStates,omitempty"`
	// The IDs of the route tables.
	RouteTableIds *[]string `json:"RouteTableIds,omitempty"`
	// The IDs of the VMs specified in routes in the tables.
	RouteVmIds *[]string `json:"RouteVmIds,omitempty"`
	// The keys of the tags associated with the route tables.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the route tables.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the route tables, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
}

// NewFiltersRouteTable instantiates a new FiltersRouteTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersRouteTable() *FiltersRouteTable {
	this := FiltersRouteTable{}
	return &this
}

// NewFiltersRouteTableWithDefaults instantiates a new FiltersRouteTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersRouteTableWithDefaults() *FiltersRouteTable {
	this := FiltersRouteTable{}
	return &this
}

// GetLinkRouteTableIds returns the LinkRouteTableIds field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetLinkRouteTableIds() []string {
	if o == nil || o.LinkRouteTableIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkRouteTableIds
}

// GetLinkRouteTableIdsOk returns a tuple with the LinkRouteTableIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetLinkRouteTableIdsOk() (*[]string, bool) {
	if o == nil || o.LinkRouteTableIds == nil {
		return nil, false
	}
	return o.LinkRouteTableIds, true
}

// HasLinkRouteTableIds returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasLinkRouteTableIds() bool {
	if o != nil && o.LinkRouteTableIds != nil {
		return true
	}

	return false
}

// SetLinkRouteTableIds gets a reference to the given []string and assigns it to the LinkRouteTableIds field.
func (o *FiltersRouteTable) SetLinkRouteTableIds(v []string) {
	o.LinkRouteTableIds = &v
}

// GetLinkRouteTableLinkRouteTableIds returns the LinkRouteTableLinkRouteTableIds field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetLinkRouteTableLinkRouteTableIds() []string {
	if o == nil || o.LinkRouteTableLinkRouteTableIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkRouteTableLinkRouteTableIds
}

// GetLinkRouteTableLinkRouteTableIdsOk returns a tuple with the LinkRouteTableLinkRouteTableIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetLinkRouteTableLinkRouteTableIdsOk() (*[]string, bool) {
	if o == nil || o.LinkRouteTableLinkRouteTableIds == nil {
		return nil, false
	}
	return o.LinkRouteTableLinkRouteTableIds, true
}

// HasLinkRouteTableLinkRouteTableIds returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasLinkRouteTableLinkRouteTableIds() bool {
	if o != nil && o.LinkRouteTableLinkRouteTableIds != nil {
		return true
	}

	return false
}

// SetLinkRouteTableLinkRouteTableIds gets a reference to the given []string and assigns it to the LinkRouteTableLinkRouteTableIds field.
func (o *FiltersRouteTable) SetLinkRouteTableLinkRouteTableIds(v []string) {
	o.LinkRouteTableLinkRouteTableIds = &v
}

// GetLinkRouteTableMain returns the LinkRouteTableMain field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetLinkRouteTableMain() bool {
	if o == nil || o.LinkRouteTableMain == nil {
		var ret bool
		return ret
	}
	return *o.LinkRouteTableMain
}

// GetLinkRouteTableMainOk returns a tuple with the LinkRouteTableMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetLinkRouteTableMainOk() (*bool, bool) {
	if o == nil || o.LinkRouteTableMain == nil {
		return nil, false
	}
	return o.LinkRouteTableMain, true
}

// HasLinkRouteTableMain returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasLinkRouteTableMain() bool {
	if o != nil && o.LinkRouteTableMain != nil {
		return true
	}

	return false
}

// SetLinkRouteTableMain gets a reference to the given bool and assigns it to the LinkRouteTableMain field.
func (o *FiltersRouteTable) SetLinkRouteTableMain(v bool) {
	o.LinkRouteTableMain = &v
}

// GetLinkSubnetIds returns the LinkSubnetIds field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetLinkSubnetIds() []string {
	if o == nil || o.LinkSubnetIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkSubnetIds
}

// GetLinkSubnetIdsOk returns a tuple with the LinkSubnetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetLinkSubnetIdsOk() (*[]string, bool) {
	if o == nil || o.LinkSubnetIds == nil {
		return nil, false
	}
	return o.LinkSubnetIds, true
}

// HasLinkSubnetIds returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasLinkSubnetIds() bool {
	if o != nil && o.LinkSubnetIds != nil {
		return true
	}

	return false
}

// SetLinkSubnetIds gets a reference to the given []string and assigns it to the LinkSubnetIds field.
func (o *FiltersRouteTable) SetLinkSubnetIds(v []string) {
	o.LinkSubnetIds = &v
}

// GetNetIds returns the NetIds field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetNetIds() []string {
	if o == nil || o.NetIds == nil {
		var ret []string
		return ret
	}
	return *o.NetIds
}

// GetNetIdsOk returns a tuple with the NetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetNetIdsOk() (*[]string, bool) {
	if o == nil || o.NetIds == nil {
		return nil, false
	}
	return o.NetIds, true
}

// HasNetIds returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasNetIds() bool {
	if o != nil && o.NetIds != nil {
		return true
	}

	return false
}

// SetNetIds gets a reference to the given []string and assigns it to the NetIds field.
func (o *FiltersRouteTable) SetNetIds(v []string) {
	o.NetIds = &v
}

// GetRouteCreationMethods returns the RouteCreationMethods field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetRouteCreationMethods() []string {
	if o == nil || o.RouteCreationMethods == nil {
		var ret []string
		return ret
	}
	return *o.RouteCreationMethods
}

// GetRouteCreationMethodsOk returns a tuple with the RouteCreationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetRouteCreationMethodsOk() (*[]string, bool) {
	if o == nil || o.RouteCreationMethods == nil {
		return nil, false
	}
	return o.RouteCreationMethods, true
}

// HasRouteCreationMethods returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasRouteCreationMethods() bool {
	if o != nil && o.RouteCreationMethods != nil {
		return true
	}

	return false
}

// SetRouteCreationMethods gets a reference to the given []string and assigns it to the RouteCreationMethods field.
func (o *FiltersRouteTable) SetRouteCreationMethods(v []string) {
	o.RouteCreationMethods = &v
}

// GetRouteDestinationIpRanges returns the RouteDestinationIpRanges field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetRouteDestinationIpRanges() []string {
	if o == nil || o.RouteDestinationIpRanges == nil {
		var ret []string
		return ret
	}
	return *o.RouteDestinationIpRanges
}

// GetRouteDestinationIpRangesOk returns a tuple with the RouteDestinationIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetRouteDestinationIpRangesOk() (*[]string, bool) {
	if o == nil || o.RouteDestinationIpRanges == nil {
		return nil, false
	}
	return o.RouteDestinationIpRanges, true
}

// HasRouteDestinationIpRanges returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasRouteDestinationIpRanges() bool {
	if o != nil && o.RouteDestinationIpRanges != nil {
		return true
	}

	return false
}

// SetRouteDestinationIpRanges gets a reference to the given []string and assigns it to the RouteDestinationIpRanges field.
func (o *FiltersRouteTable) SetRouteDestinationIpRanges(v []string) {
	o.RouteDestinationIpRanges = &v
}

// GetRouteDestinationServiceIds returns the RouteDestinationServiceIds field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetRouteDestinationServiceIds() []string {
	if o == nil || o.RouteDestinationServiceIds == nil {
		var ret []string
		return ret
	}
	return *o.RouteDestinationServiceIds
}

// GetRouteDestinationServiceIdsOk returns a tuple with the RouteDestinationServiceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetRouteDestinationServiceIdsOk() (*[]string, bool) {
	if o == nil || o.RouteDestinationServiceIds == nil {
		return nil, false
	}
	return o.RouteDestinationServiceIds, true
}

// HasRouteDestinationServiceIds returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasRouteDestinationServiceIds() bool {
	if o != nil && o.RouteDestinationServiceIds != nil {
		return true
	}

	return false
}

// SetRouteDestinationServiceIds gets a reference to the given []string and assigns it to the RouteDestinationServiceIds field.
func (o *FiltersRouteTable) SetRouteDestinationServiceIds(v []string) {
	o.RouteDestinationServiceIds = &v
}

// GetRouteGatewayIds returns the RouteGatewayIds field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetRouteGatewayIds() []string {
	if o == nil || o.RouteGatewayIds == nil {
		var ret []string
		return ret
	}
	return *o.RouteGatewayIds
}

// GetRouteGatewayIdsOk returns a tuple with the RouteGatewayIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetRouteGatewayIdsOk() (*[]string, bool) {
	if o == nil || o.RouteGatewayIds == nil {
		return nil, false
	}
	return o.RouteGatewayIds, true
}

// HasRouteGatewayIds returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasRouteGatewayIds() bool {
	if o != nil && o.RouteGatewayIds != nil {
		return true
	}

	return false
}

// SetRouteGatewayIds gets a reference to the given []string and assigns it to the RouteGatewayIds field.
func (o *FiltersRouteTable) SetRouteGatewayIds(v []string) {
	o.RouteGatewayIds = &v
}

// GetRouteNatServiceIds returns the RouteNatServiceIds field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetRouteNatServiceIds() []string {
	if o == nil || o.RouteNatServiceIds == nil {
		var ret []string
		return ret
	}
	return *o.RouteNatServiceIds
}

// GetRouteNatServiceIdsOk returns a tuple with the RouteNatServiceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetRouteNatServiceIdsOk() (*[]string, bool) {
	if o == nil || o.RouteNatServiceIds == nil {
		return nil, false
	}
	return o.RouteNatServiceIds, true
}

// HasRouteNatServiceIds returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasRouteNatServiceIds() bool {
	if o != nil && o.RouteNatServiceIds != nil {
		return true
	}

	return false
}

// SetRouteNatServiceIds gets a reference to the given []string and assigns it to the RouteNatServiceIds field.
func (o *FiltersRouteTable) SetRouteNatServiceIds(v []string) {
	o.RouteNatServiceIds = &v
}

// GetRouteNetPeeringIds returns the RouteNetPeeringIds field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetRouteNetPeeringIds() []string {
	if o == nil || o.RouteNetPeeringIds == nil {
		var ret []string
		return ret
	}
	return *o.RouteNetPeeringIds
}

// GetRouteNetPeeringIdsOk returns a tuple with the RouteNetPeeringIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetRouteNetPeeringIdsOk() (*[]string, bool) {
	if o == nil || o.RouteNetPeeringIds == nil {
		return nil, false
	}
	return o.RouteNetPeeringIds, true
}

// HasRouteNetPeeringIds returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasRouteNetPeeringIds() bool {
	if o != nil && o.RouteNetPeeringIds != nil {
		return true
	}

	return false
}

// SetRouteNetPeeringIds gets a reference to the given []string and assigns it to the RouteNetPeeringIds field.
func (o *FiltersRouteTable) SetRouteNetPeeringIds(v []string) {
	o.RouteNetPeeringIds = &v
}

// GetRouteStates returns the RouteStates field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetRouteStates() []string {
	if o == nil || o.RouteStates == nil {
		var ret []string
		return ret
	}
	return *o.RouteStates
}

// GetRouteStatesOk returns a tuple with the RouteStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetRouteStatesOk() (*[]string, bool) {
	if o == nil || o.RouteStates == nil {
		return nil, false
	}
	return o.RouteStates, true
}

// HasRouteStates returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasRouteStates() bool {
	if o != nil && o.RouteStates != nil {
		return true
	}

	return false
}

// SetRouteStates gets a reference to the given []string and assigns it to the RouteStates field.
func (o *FiltersRouteTable) SetRouteStates(v []string) {
	o.RouteStates = &v
}

// GetRouteTableIds returns the RouteTableIds field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetRouteTableIds() []string {
	if o == nil || o.RouteTableIds == nil {
		var ret []string
		return ret
	}
	return *o.RouteTableIds
}

// GetRouteTableIdsOk returns a tuple with the RouteTableIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetRouteTableIdsOk() (*[]string, bool) {
	if o == nil || o.RouteTableIds == nil {
		return nil, false
	}
	return o.RouteTableIds, true
}

// HasRouteTableIds returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasRouteTableIds() bool {
	if o != nil && o.RouteTableIds != nil {
		return true
	}

	return false
}

// SetRouteTableIds gets a reference to the given []string and assigns it to the RouteTableIds field.
func (o *FiltersRouteTable) SetRouteTableIds(v []string) {
	o.RouteTableIds = &v
}

// GetRouteVmIds returns the RouteVmIds field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetRouteVmIds() []string {
	if o == nil || o.RouteVmIds == nil {
		var ret []string
		return ret
	}
	return *o.RouteVmIds
}

// GetRouteVmIdsOk returns a tuple with the RouteVmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetRouteVmIdsOk() (*[]string, bool) {
	if o == nil || o.RouteVmIds == nil {
		return nil, false
	}
	return o.RouteVmIds, true
}

// HasRouteVmIds returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasRouteVmIds() bool {
	if o != nil && o.RouteVmIds != nil {
		return true
	}

	return false
}

// SetRouteVmIds gets a reference to the given []string and assigns it to the RouteVmIds field.
func (o *FiltersRouteTable) SetRouteVmIds(v []string) {
	o.RouteVmIds = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersRouteTable) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersRouteTable) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersRouteTable) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersRouteTable) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersRouteTable) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersRouteTable) SetTags(v []string) {
	o.Tags = &v
}

func (o FiltersRouteTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinkRouteTableIds != nil {
		toSerialize["LinkRouteTableIds"] = o.LinkRouteTableIds
	}
	if o.LinkRouteTableLinkRouteTableIds != nil {
		toSerialize["LinkRouteTableLinkRouteTableIds"] = o.LinkRouteTableLinkRouteTableIds
	}
	if o.LinkRouteTableMain != nil {
		toSerialize["LinkRouteTableMain"] = o.LinkRouteTableMain
	}
	if o.LinkSubnetIds != nil {
		toSerialize["LinkSubnetIds"] = o.LinkSubnetIds
	}
	if o.NetIds != nil {
		toSerialize["NetIds"] = o.NetIds
	}
	if o.RouteCreationMethods != nil {
		toSerialize["RouteCreationMethods"] = o.RouteCreationMethods
	}
	if o.RouteDestinationIpRanges != nil {
		toSerialize["RouteDestinationIpRanges"] = o.RouteDestinationIpRanges
	}
	if o.RouteDestinationServiceIds != nil {
		toSerialize["RouteDestinationServiceIds"] = o.RouteDestinationServiceIds
	}
	if o.RouteGatewayIds != nil {
		toSerialize["RouteGatewayIds"] = o.RouteGatewayIds
	}
	if o.RouteNatServiceIds != nil {
		toSerialize["RouteNatServiceIds"] = o.RouteNatServiceIds
	}
	if o.RouteNetPeeringIds != nil {
		toSerialize["RouteNetPeeringIds"] = o.RouteNetPeeringIds
	}
	if o.RouteStates != nil {
		toSerialize["RouteStates"] = o.RouteStates
	}
	if o.RouteTableIds != nil {
		toSerialize["RouteTableIds"] = o.RouteTableIds
	}
	if o.RouteVmIds != nil {
		toSerialize["RouteVmIds"] = o.RouteVmIds
	}
	if o.TagKeys != nil {
		toSerialize["TagKeys"] = o.TagKeys
	}
	if o.TagValues != nil {
		toSerialize["TagValues"] = o.TagValues
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersRouteTable struct {
	value *FiltersRouteTable
	isSet bool
}

func (v NullableFiltersRouteTable) Get() *FiltersRouteTable {
	return v.value
}

func (v *NullableFiltersRouteTable) Set(val *FiltersRouteTable) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersRouteTable) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersRouteTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersRouteTable(val *FiltersRouteTable) *NullableFiltersRouteTable {
	return &NullableFiltersRouteTable{value: val, isSet: true}
}

func (v NullableFiltersRouteTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersRouteTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
