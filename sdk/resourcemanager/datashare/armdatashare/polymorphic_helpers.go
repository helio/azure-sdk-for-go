//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatashare

import "encoding/json"

func unmarshalDataSetClassification(rawMsg json.RawMessage) (DataSetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataSetClassification
	switch m["kind"] {
	case string(DataSetKindAdlsGen1File):
		b = &ADLSGen1FileDataSet{}
	case string(DataSetKindAdlsGen1Folder):
		b = &ADLSGen1FolderDataSet{}
	case string(DataSetKindAdlsGen2File):
		b = &ADLSGen2FileDataSet{}
	case string(DataSetKindAdlsGen2FileSystem):
		b = &ADLSGen2FileSystemDataSet{}
	case string(DataSetKindAdlsGen2Folder):
		b = &ADLSGen2FolderDataSet{}
	case string(DataSetKindBlob):
		b = &BlobDataSet{}
	case string(DataSetKindBlobFolder):
		b = &BlobFolderDataSet{}
	case string(DataSetKindContainer):
		b = &BlobContainerDataSet{}
	case string(DataSetKindKustoCluster):
		b = &KustoClusterDataSet{}
	case string(DataSetKindKustoDatabase):
		b = &KustoDatabaseDataSet{}
	case string(DataSetKindSQLDBTable):
		b = &SQLDBTableDataSet{}
	case string(DataSetKindSQLDWTable):
		b = &SQLDWTableDataSet{}
	case string(DataSetKindSynapseWorkspaceSQLPoolTable):
		b = &SynapseWorkspaceSQLPoolTableDataSet{}
	default:
		b = &DataSet{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDataSetClassificationArray(rawMsg json.RawMessage) ([]DataSetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DataSetClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDataSetClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDataSetMappingClassification(rawMsg json.RawMessage) (DataSetMappingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataSetMappingClassification
	switch m["kind"] {
	case string(DataSetMappingKindAdlsGen2File):
		b = &ADLSGen2FileDataSetMapping{}
	case string(DataSetMappingKindAdlsGen2FileSystem):
		b = &ADLSGen2FileSystemDataSetMapping{}
	case string(DataSetMappingKindAdlsGen2Folder):
		b = &ADLSGen2FolderDataSetMapping{}
	case string(DataSetMappingKindBlob):
		b = &BlobDataSetMapping{}
	case string(DataSetMappingKindBlobFolder):
		b = &BlobFolderDataSetMapping{}
	case string(DataSetMappingKindContainer):
		b = &BlobContainerDataSetMapping{}
	case string(DataSetMappingKindKustoCluster):
		b = &KustoClusterDataSetMapping{}
	case string(DataSetMappingKindKustoDatabase):
		b = &KustoDatabaseDataSetMapping{}
	case string(DataSetMappingKindSQLDBTable):
		b = &SQLDBTableDataSetMapping{}
	case string(DataSetMappingKindSQLDWTable):
		b = &SQLDWTableDataSetMapping{}
	case string(DataSetMappingKindSynapseWorkspaceSQLPoolTable):
		b = &SynapseWorkspaceSQLPoolTableDataSetMapping{}
	default:
		b = &DataSetMapping{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDataSetMappingClassificationArray(rawMsg json.RawMessage) ([]DataSetMappingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DataSetMappingClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDataSetMappingClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalSourceShareSynchronizationSettingClassification(rawMsg json.RawMessage) (SourceShareSynchronizationSettingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SourceShareSynchronizationSettingClassification
	switch m["kind"] {
	case string(SourceShareSynchronizationSettingKindScheduleBased):
		b = &ScheduledSourceSynchronizationSetting{}
	default:
		b = &SourceShareSynchronizationSetting{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSourceShareSynchronizationSettingClassificationArray(rawMsg json.RawMessage) ([]SourceShareSynchronizationSettingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]SourceShareSynchronizationSettingClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalSourceShareSynchronizationSettingClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalSynchronizationSettingClassification(rawMsg json.RawMessage) (SynchronizationSettingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SynchronizationSettingClassification
	switch m["kind"] {
	case string(SynchronizationSettingKindScheduleBased):
		b = &ScheduledSynchronizationSetting{}
	default:
		b = &SynchronizationSetting{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSynchronizationSettingClassificationArray(rawMsg json.RawMessage) ([]SynchronizationSettingClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]SynchronizationSettingClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalSynchronizationSettingClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalTriggerClassification(rawMsg json.RawMessage) (TriggerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TriggerClassification
	switch m["kind"] {
	case string(TriggerKindScheduleBased):
		b = &ScheduledTrigger{}
	default:
		b = &Trigger{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTriggerClassificationArray(rawMsg json.RawMessage) ([]TriggerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]TriggerClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalTriggerClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}