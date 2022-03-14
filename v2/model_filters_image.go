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

// FiltersImage One or more filters.
type FiltersImage struct {
	// The account aliases of the owners of the OMIs.
	AccountAliases *[]string `json:"AccountAliases,omitempty"`
	// The account IDs of the owners of the OMIs. By default, all the OMIs for which you have launch permissions are described.
	AccountIds *[]string `json:"AccountIds,omitempty"`
	// The architectures of the OMIs (`i386` \\| `x86_64`).
	Architectures *[]string `json:"Architectures,omitempty"`
	// Whether the volumes are deleted or not when terminating the VM.
	BlockDeviceMappingDeleteOnVmDeletion *bool `json:"BlockDeviceMappingDeleteOnVmDeletion,omitempty"`
	// The device names for the volumes.
	BlockDeviceMappingDeviceNames *[]string `json:"BlockDeviceMappingDeviceNames,omitempty"`
	// The IDs of the snapshots used to create the volumes.
	BlockDeviceMappingSnapshotIds *[]string `json:"BlockDeviceMappingSnapshotIds,omitempty"`
	// The sizes of the volumes, in gibibytes (GiB).
	BlockDeviceMappingVolumeSizes *[]int32 `json:"BlockDeviceMappingVolumeSizes,omitempty"`
	// The types of volumes (`standard` \\| `gp2` \\| `io1`).
	BlockDeviceMappingVolumeTypes *[]string `json:"BlockDeviceMappingVolumeTypes,omitempty"`
	// The descriptions of the OMIs, provided when they were created.
	Descriptions *[]string `json:"Descriptions,omitempty"`
	// The locations of the buckets where the OMI files are stored.
	FileLocations *[]string `json:"FileLocations,omitempty"`
	// The hypervisor type of the OMI (always `xen`).
	Hypervisors *[]string `json:"Hypervisors,omitempty"`
	// The IDs of the OMIs.
	ImageIds *[]string `json:"ImageIds,omitempty"`
	// The names of the OMIs, provided when they were created.
	ImageNames *[]string `json:"ImageNames,omitempty"`
	// The account IDs of the users who have launch permissions for the OMIs.
	PermissionsToLaunchAccountIds *[]string `json:"PermissionsToLaunchAccountIds,omitempty"`
	// If true, lists all public OMIs. If false, lists all private OMIs.
	PermissionsToLaunchGlobalPermission *bool `json:"PermissionsToLaunchGlobalPermission,omitempty"`
	// The product code associated with the OMI (`0001` Linux/Unix \\| `0002` Windows \\| `0004` Linux/Oracle \\| `0005` Windows 10).
	ProductCodes *[]string `json:"ProductCodes,omitempty"`
	// The device names of the root devices (for example, `/dev/sda1`).
	RootDeviceNames *[]string `json:"RootDeviceNames,omitempty"`
	// The types of root device used by the OMIs (always `bsu`).
	RootDeviceTypes *[]string `json:"RootDeviceTypes,omitempty"`
	// The states of the OMIs (`pending` \\| `available` \\| `failed`).
	States *[]string `json:"States,omitempty"`
	// The keys of the tags associated with the OMIs.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the OMIs.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the OMIs, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
	// The virtualization types (always `hvm`).
	VirtualizationTypes *[]string `json:"VirtualizationTypes,omitempty"`
}

// NewFiltersImage instantiates a new FiltersImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersImage() *FiltersImage {
	this := FiltersImage{}
	return &this
}

// NewFiltersImageWithDefaults instantiates a new FiltersImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersImageWithDefaults() *FiltersImage {
	this := FiltersImage{}
	return &this
}

// GetAccountAliases returns the AccountAliases field value if set, zero value otherwise.
func (o *FiltersImage) GetAccountAliases() []string {
	if o == nil || o.AccountAliases == nil {
		var ret []string
		return ret
	}
	return *o.AccountAliases
}

// GetAccountAliasesOk returns a tuple with the AccountAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetAccountAliasesOk() (*[]string, bool) {
	if o == nil || o.AccountAliases == nil {
		return nil, false
	}
	return o.AccountAliases, true
}

// HasAccountAliases returns a boolean if a field has been set.
func (o *FiltersImage) HasAccountAliases() bool {
	if o != nil && o.AccountAliases != nil {
		return true
	}

	return false
}

// SetAccountAliases gets a reference to the given []string and assigns it to the AccountAliases field.
func (o *FiltersImage) SetAccountAliases(v []string) {
	o.AccountAliases = &v
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *FiltersImage) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *FiltersImage) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *FiltersImage) SetAccountIds(v []string) {
	o.AccountIds = &v
}

// GetArchitectures returns the Architectures field value if set, zero value otherwise.
func (o *FiltersImage) GetArchitectures() []string {
	if o == nil || o.Architectures == nil {
		var ret []string
		return ret
	}
	return *o.Architectures
}

// GetArchitecturesOk returns a tuple with the Architectures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetArchitecturesOk() (*[]string, bool) {
	if o == nil || o.Architectures == nil {
		return nil, false
	}
	return o.Architectures, true
}

// HasArchitectures returns a boolean if a field has been set.
func (o *FiltersImage) HasArchitectures() bool {
	if o != nil && o.Architectures != nil {
		return true
	}

	return false
}

// SetArchitectures gets a reference to the given []string and assigns it to the Architectures field.
func (o *FiltersImage) SetArchitectures(v []string) {
	o.Architectures = &v
}

// GetBlockDeviceMappingDeleteOnVmDeletion returns the BlockDeviceMappingDeleteOnVmDeletion field value if set, zero value otherwise.
func (o *FiltersImage) GetBlockDeviceMappingDeleteOnVmDeletion() bool {
	if o == nil || o.BlockDeviceMappingDeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.BlockDeviceMappingDeleteOnVmDeletion
}

// GetBlockDeviceMappingDeleteOnVmDeletionOk returns a tuple with the BlockDeviceMappingDeleteOnVmDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetBlockDeviceMappingDeleteOnVmDeletionOk() (*bool, bool) {
	if o == nil || o.BlockDeviceMappingDeleteOnVmDeletion == nil {
		return nil, false
	}
	return o.BlockDeviceMappingDeleteOnVmDeletion, true
}

// HasBlockDeviceMappingDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *FiltersImage) HasBlockDeviceMappingDeleteOnVmDeletion() bool {
	if o != nil && o.BlockDeviceMappingDeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetBlockDeviceMappingDeleteOnVmDeletion gets a reference to the given bool and assigns it to the BlockDeviceMappingDeleteOnVmDeletion field.
func (o *FiltersImage) SetBlockDeviceMappingDeleteOnVmDeletion(v bool) {
	o.BlockDeviceMappingDeleteOnVmDeletion = &v
}

// GetBlockDeviceMappingDeviceNames returns the BlockDeviceMappingDeviceNames field value if set, zero value otherwise.
func (o *FiltersImage) GetBlockDeviceMappingDeviceNames() []string {
	if o == nil || o.BlockDeviceMappingDeviceNames == nil {
		var ret []string
		return ret
	}
	return *o.BlockDeviceMappingDeviceNames
}

// GetBlockDeviceMappingDeviceNamesOk returns a tuple with the BlockDeviceMappingDeviceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetBlockDeviceMappingDeviceNamesOk() (*[]string, bool) {
	if o == nil || o.BlockDeviceMappingDeviceNames == nil {
		return nil, false
	}
	return o.BlockDeviceMappingDeviceNames, true
}

// HasBlockDeviceMappingDeviceNames returns a boolean if a field has been set.
func (o *FiltersImage) HasBlockDeviceMappingDeviceNames() bool {
	if o != nil && o.BlockDeviceMappingDeviceNames != nil {
		return true
	}

	return false
}

// SetBlockDeviceMappingDeviceNames gets a reference to the given []string and assigns it to the BlockDeviceMappingDeviceNames field.
func (o *FiltersImage) SetBlockDeviceMappingDeviceNames(v []string) {
	o.BlockDeviceMappingDeviceNames = &v
}

// GetBlockDeviceMappingSnapshotIds returns the BlockDeviceMappingSnapshotIds field value if set, zero value otherwise.
func (o *FiltersImage) GetBlockDeviceMappingSnapshotIds() []string {
	if o == nil || o.BlockDeviceMappingSnapshotIds == nil {
		var ret []string
		return ret
	}
	return *o.BlockDeviceMappingSnapshotIds
}

// GetBlockDeviceMappingSnapshotIdsOk returns a tuple with the BlockDeviceMappingSnapshotIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetBlockDeviceMappingSnapshotIdsOk() (*[]string, bool) {
	if o == nil || o.BlockDeviceMappingSnapshotIds == nil {
		return nil, false
	}
	return o.BlockDeviceMappingSnapshotIds, true
}

// HasBlockDeviceMappingSnapshotIds returns a boolean if a field has been set.
func (o *FiltersImage) HasBlockDeviceMappingSnapshotIds() bool {
	if o != nil && o.BlockDeviceMappingSnapshotIds != nil {
		return true
	}

	return false
}

// SetBlockDeviceMappingSnapshotIds gets a reference to the given []string and assigns it to the BlockDeviceMappingSnapshotIds field.
func (o *FiltersImage) SetBlockDeviceMappingSnapshotIds(v []string) {
	o.BlockDeviceMappingSnapshotIds = &v
}

// GetBlockDeviceMappingVolumeSizes returns the BlockDeviceMappingVolumeSizes field value if set, zero value otherwise.
func (o *FiltersImage) GetBlockDeviceMappingVolumeSizes() []int32 {
	if o == nil || o.BlockDeviceMappingVolumeSizes == nil {
		var ret []int32
		return ret
	}
	return *o.BlockDeviceMappingVolumeSizes
}

// GetBlockDeviceMappingVolumeSizesOk returns a tuple with the BlockDeviceMappingVolumeSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetBlockDeviceMappingVolumeSizesOk() (*[]int32, bool) {
	if o == nil || o.BlockDeviceMappingVolumeSizes == nil {
		return nil, false
	}
	return o.BlockDeviceMappingVolumeSizes, true
}

// HasBlockDeviceMappingVolumeSizes returns a boolean if a field has been set.
func (o *FiltersImage) HasBlockDeviceMappingVolumeSizes() bool {
	if o != nil && o.BlockDeviceMappingVolumeSizes != nil {
		return true
	}

	return false
}

// SetBlockDeviceMappingVolumeSizes gets a reference to the given []int32 and assigns it to the BlockDeviceMappingVolumeSizes field.
func (o *FiltersImage) SetBlockDeviceMappingVolumeSizes(v []int32) {
	o.BlockDeviceMappingVolumeSizes = &v
}

// GetBlockDeviceMappingVolumeTypes returns the BlockDeviceMappingVolumeTypes field value if set, zero value otherwise.
func (o *FiltersImage) GetBlockDeviceMappingVolumeTypes() []string {
	if o == nil || o.BlockDeviceMappingVolumeTypes == nil {
		var ret []string
		return ret
	}
	return *o.BlockDeviceMappingVolumeTypes
}

// GetBlockDeviceMappingVolumeTypesOk returns a tuple with the BlockDeviceMappingVolumeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetBlockDeviceMappingVolumeTypesOk() (*[]string, bool) {
	if o == nil || o.BlockDeviceMappingVolumeTypes == nil {
		return nil, false
	}
	return o.BlockDeviceMappingVolumeTypes, true
}

// HasBlockDeviceMappingVolumeTypes returns a boolean if a field has been set.
func (o *FiltersImage) HasBlockDeviceMappingVolumeTypes() bool {
	if o != nil && o.BlockDeviceMappingVolumeTypes != nil {
		return true
	}

	return false
}

// SetBlockDeviceMappingVolumeTypes gets a reference to the given []string and assigns it to the BlockDeviceMappingVolumeTypes field.
func (o *FiltersImage) SetBlockDeviceMappingVolumeTypes(v []string) {
	o.BlockDeviceMappingVolumeTypes = &v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *FiltersImage) GetDescriptions() []string {
	if o == nil || o.Descriptions == nil {
		var ret []string
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetDescriptionsOk() (*[]string, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *FiltersImage) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []string and assigns it to the Descriptions field.
func (o *FiltersImage) SetDescriptions(v []string) {
	o.Descriptions = &v
}

// GetFileLocations returns the FileLocations field value if set, zero value otherwise.
func (o *FiltersImage) GetFileLocations() []string {
	if o == nil || o.FileLocations == nil {
		var ret []string
		return ret
	}
	return *o.FileLocations
}

// GetFileLocationsOk returns a tuple with the FileLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetFileLocationsOk() (*[]string, bool) {
	if o == nil || o.FileLocations == nil {
		return nil, false
	}
	return o.FileLocations, true
}

// HasFileLocations returns a boolean if a field has been set.
func (o *FiltersImage) HasFileLocations() bool {
	if o != nil && o.FileLocations != nil {
		return true
	}

	return false
}

// SetFileLocations gets a reference to the given []string and assigns it to the FileLocations field.
func (o *FiltersImage) SetFileLocations(v []string) {
	o.FileLocations = &v
}

// GetHypervisors returns the Hypervisors field value if set, zero value otherwise.
func (o *FiltersImage) GetHypervisors() []string {
	if o == nil || o.Hypervisors == nil {
		var ret []string
		return ret
	}
	return *o.Hypervisors
}

// GetHypervisorsOk returns a tuple with the Hypervisors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetHypervisorsOk() (*[]string, bool) {
	if o == nil || o.Hypervisors == nil {
		return nil, false
	}
	return o.Hypervisors, true
}

// HasHypervisors returns a boolean if a field has been set.
func (o *FiltersImage) HasHypervisors() bool {
	if o != nil && o.Hypervisors != nil {
		return true
	}

	return false
}

// SetHypervisors gets a reference to the given []string and assigns it to the Hypervisors field.
func (o *FiltersImage) SetHypervisors(v []string) {
	o.Hypervisors = &v
}

// GetImageIds returns the ImageIds field value if set, zero value otherwise.
func (o *FiltersImage) GetImageIds() []string {
	if o == nil || o.ImageIds == nil {
		var ret []string
		return ret
	}
	return *o.ImageIds
}

// GetImageIdsOk returns a tuple with the ImageIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetImageIdsOk() (*[]string, bool) {
	if o == nil || o.ImageIds == nil {
		return nil, false
	}
	return o.ImageIds, true
}

// HasImageIds returns a boolean if a field has been set.
func (o *FiltersImage) HasImageIds() bool {
	if o != nil && o.ImageIds != nil {
		return true
	}

	return false
}

// SetImageIds gets a reference to the given []string and assigns it to the ImageIds field.
func (o *FiltersImage) SetImageIds(v []string) {
	o.ImageIds = &v
}

// GetImageNames returns the ImageNames field value if set, zero value otherwise.
func (o *FiltersImage) GetImageNames() []string {
	if o == nil || o.ImageNames == nil {
		var ret []string
		return ret
	}
	return *o.ImageNames
}

// GetImageNamesOk returns a tuple with the ImageNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetImageNamesOk() (*[]string, bool) {
	if o == nil || o.ImageNames == nil {
		return nil, false
	}
	return o.ImageNames, true
}

// HasImageNames returns a boolean if a field has been set.
func (o *FiltersImage) HasImageNames() bool {
	if o != nil && o.ImageNames != nil {
		return true
	}

	return false
}

// SetImageNames gets a reference to the given []string and assigns it to the ImageNames field.
func (o *FiltersImage) SetImageNames(v []string) {
	o.ImageNames = &v
}

// GetPermissionsToLaunchAccountIds returns the PermissionsToLaunchAccountIds field value if set, zero value otherwise.
func (o *FiltersImage) GetPermissionsToLaunchAccountIds() []string {
	if o == nil || o.PermissionsToLaunchAccountIds == nil {
		var ret []string
		return ret
	}
	return *o.PermissionsToLaunchAccountIds
}

// GetPermissionsToLaunchAccountIdsOk returns a tuple with the PermissionsToLaunchAccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetPermissionsToLaunchAccountIdsOk() (*[]string, bool) {
	if o == nil || o.PermissionsToLaunchAccountIds == nil {
		return nil, false
	}
	return o.PermissionsToLaunchAccountIds, true
}

// HasPermissionsToLaunchAccountIds returns a boolean if a field has been set.
func (o *FiltersImage) HasPermissionsToLaunchAccountIds() bool {
	if o != nil && o.PermissionsToLaunchAccountIds != nil {
		return true
	}

	return false
}

// SetPermissionsToLaunchAccountIds gets a reference to the given []string and assigns it to the PermissionsToLaunchAccountIds field.
func (o *FiltersImage) SetPermissionsToLaunchAccountIds(v []string) {
	o.PermissionsToLaunchAccountIds = &v
}

// GetPermissionsToLaunchGlobalPermission returns the PermissionsToLaunchGlobalPermission field value if set, zero value otherwise.
func (o *FiltersImage) GetPermissionsToLaunchGlobalPermission() bool {
	if o == nil || o.PermissionsToLaunchGlobalPermission == nil {
		var ret bool
		return ret
	}
	return *o.PermissionsToLaunchGlobalPermission
}

// GetPermissionsToLaunchGlobalPermissionOk returns a tuple with the PermissionsToLaunchGlobalPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetPermissionsToLaunchGlobalPermissionOk() (*bool, bool) {
	if o == nil || o.PermissionsToLaunchGlobalPermission == nil {
		return nil, false
	}
	return o.PermissionsToLaunchGlobalPermission, true
}

// HasPermissionsToLaunchGlobalPermission returns a boolean if a field has been set.
func (o *FiltersImage) HasPermissionsToLaunchGlobalPermission() bool {
	if o != nil && o.PermissionsToLaunchGlobalPermission != nil {
		return true
	}

	return false
}

// SetPermissionsToLaunchGlobalPermission gets a reference to the given bool and assigns it to the PermissionsToLaunchGlobalPermission field.
func (o *FiltersImage) SetPermissionsToLaunchGlobalPermission(v bool) {
	o.PermissionsToLaunchGlobalPermission = &v
}

// GetProductCodes returns the ProductCodes field value if set, zero value otherwise.
func (o *FiltersImage) GetProductCodes() []string {
	if o == nil || o.ProductCodes == nil {
		var ret []string
		return ret
	}
	return *o.ProductCodes
}

// GetProductCodesOk returns a tuple with the ProductCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetProductCodesOk() (*[]string, bool) {
	if o == nil || o.ProductCodes == nil {
		return nil, false
	}
	return o.ProductCodes, true
}

// HasProductCodes returns a boolean if a field has been set.
func (o *FiltersImage) HasProductCodes() bool {
	if o != nil && o.ProductCodes != nil {
		return true
	}

	return false
}

// SetProductCodes gets a reference to the given []string and assigns it to the ProductCodes field.
func (o *FiltersImage) SetProductCodes(v []string) {
	o.ProductCodes = &v
}

// GetRootDeviceNames returns the RootDeviceNames field value if set, zero value otherwise.
func (o *FiltersImage) GetRootDeviceNames() []string {
	if o == nil || o.RootDeviceNames == nil {
		var ret []string
		return ret
	}
	return *o.RootDeviceNames
}

// GetRootDeviceNamesOk returns a tuple with the RootDeviceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetRootDeviceNamesOk() (*[]string, bool) {
	if o == nil || o.RootDeviceNames == nil {
		return nil, false
	}
	return o.RootDeviceNames, true
}

// HasRootDeviceNames returns a boolean if a field has been set.
func (o *FiltersImage) HasRootDeviceNames() bool {
	if o != nil && o.RootDeviceNames != nil {
		return true
	}

	return false
}

// SetRootDeviceNames gets a reference to the given []string and assigns it to the RootDeviceNames field.
func (o *FiltersImage) SetRootDeviceNames(v []string) {
	o.RootDeviceNames = &v
}

// GetRootDeviceTypes returns the RootDeviceTypes field value if set, zero value otherwise.
func (o *FiltersImage) GetRootDeviceTypes() []string {
	if o == nil || o.RootDeviceTypes == nil {
		var ret []string
		return ret
	}
	return *o.RootDeviceTypes
}

// GetRootDeviceTypesOk returns a tuple with the RootDeviceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetRootDeviceTypesOk() (*[]string, bool) {
	if o == nil || o.RootDeviceTypes == nil {
		return nil, false
	}
	return o.RootDeviceTypes, true
}

// HasRootDeviceTypes returns a boolean if a field has been set.
func (o *FiltersImage) HasRootDeviceTypes() bool {
	if o != nil && o.RootDeviceTypes != nil {
		return true
	}

	return false
}

// SetRootDeviceTypes gets a reference to the given []string and assigns it to the RootDeviceTypes field.
func (o *FiltersImage) SetRootDeviceTypes(v []string) {
	o.RootDeviceTypes = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *FiltersImage) GetStates() []string {
	if o == nil || o.States == nil {
		var ret []string
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetStatesOk() (*[]string, bool) {
	if o == nil || o.States == nil {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *FiltersImage) HasStates() bool {
	if o != nil && o.States != nil {
		return true
	}

	return false
}

// SetStates gets a reference to the given []string and assigns it to the States field.
func (o *FiltersImage) SetStates(v []string) {
	o.States = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersImage) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersImage) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersImage) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersImage) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersImage) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersImage) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersImage) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersImage) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersImage) SetTags(v []string) {
	o.Tags = &v
}

// GetVirtualizationTypes returns the VirtualizationTypes field value if set, zero value otherwise.
func (o *FiltersImage) GetVirtualizationTypes() []string {
	if o == nil || o.VirtualizationTypes == nil {
		var ret []string
		return ret
	}
	return *o.VirtualizationTypes
}

// GetVirtualizationTypesOk returns a tuple with the VirtualizationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersImage) GetVirtualizationTypesOk() (*[]string, bool) {
	if o == nil || o.VirtualizationTypes == nil {
		return nil, false
	}
	return o.VirtualizationTypes, true
}

// HasVirtualizationTypes returns a boolean if a field has been set.
func (o *FiltersImage) HasVirtualizationTypes() bool {
	if o != nil && o.VirtualizationTypes != nil {
		return true
	}

	return false
}

// SetVirtualizationTypes gets a reference to the given []string and assigns it to the VirtualizationTypes field.
func (o *FiltersImage) SetVirtualizationTypes(v []string) {
	o.VirtualizationTypes = &v
}

func (o FiltersImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountAliases != nil {
		toSerialize["AccountAliases"] = o.AccountAliases
	}
	if o.AccountIds != nil {
		toSerialize["AccountIds"] = o.AccountIds
	}
	if o.Architectures != nil {
		toSerialize["Architectures"] = o.Architectures
	}
	if o.BlockDeviceMappingDeleteOnVmDeletion != nil {
		toSerialize["BlockDeviceMappingDeleteOnVmDeletion"] = o.BlockDeviceMappingDeleteOnVmDeletion
	}
	if o.BlockDeviceMappingDeviceNames != nil {
		toSerialize["BlockDeviceMappingDeviceNames"] = o.BlockDeviceMappingDeviceNames
	}
	if o.BlockDeviceMappingSnapshotIds != nil {
		toSerialize["BlockDeviceMappingSnapshotIds"] = o.BlockDeviceMappingSnapshotIds
	}
	if o.BlockDeviceMappingVolumeSizes != nil {
		toSerialize["BlockDeviceMappingVolumeSizes"] = o.BlockDeviceMappingVolumeSizes
	}
	if o.BlockDeviceMappingVolumeTypes != nil {
		toSerialize["BlockDeviceMappingVolumeTypes"] = o.BlockDeviceMappingVolumeTypes
	}
	if o.Descriptions != nil {
		toSerialize["Descriptions"] = o.Descriptions
	}
	if o.FileLocations != nil {
		toSerialize["FileLocations"] = o.FileLocations
	}
	if o.Hypervisors != nil {
		toSerialize["Hypervisors"] = o.Hypervisors
	}
	if o.ImageIds != nil {
		toSerialize["ImageIds"] = o.ImageIds
	}
	if o.ImageNames != nil {
		toSerialize["ImageNames"] = o.ImageNames
	}
	if o.PermissionsToLaunchAccountIds != nil {
		toSerialize["PermissionsToLaunchAccountIds"] = o.PermissionsToLaunchAccountIds
	}
	if o.PermissionsToLaunchGlobalPermission != nil {
		toSerialize["PermissionsToLaunchGlobalPermission"] = o.PermissionsToLaunchGlobalPermission
	}
	if o.ProductCodes != nil {
		toSerialize["ProductCodes"] = o.ProductCodes
	}
	if o.RootDeviceNames != nil {
		toSerialize["RootDeviceNames"] = o.RootDeviceNames
	}
	if o.RootDeviceTypes != nil {
		toSerialize["RootDeviceTypes"] = o.RootDeviceTypes
	}
	if o.States != nil {
		toSerialize["States"] = o.States
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
	if o.VirtualizationTypes != nil {
		toSerialize["VirtualizationTypes"] = o.VirtualizationTypes
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersImage struct {
	value *FiltersImage
	isSet bool
}

func (v NullableFiltersImage) Get() *FiltersImage {
	return v.value
}

func (v *NullableFiltersImage) Set(val *FiltersImage) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersImage) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersImage(val *FiltersImage) *NullableFiltersImage {
	return &NullableFiltersImage{value: val, isSet: true}
}

func (v NullableFiltersImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
