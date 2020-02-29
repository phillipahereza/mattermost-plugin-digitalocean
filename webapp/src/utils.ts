/* eslint-disable @typescript-eslint/explicit-function-return-type */
import {changeOpacity} from 'mattermost-redux/utils/theme_utils';

import {GenericSelectData} from './ts_types';

export function prepareRegionsSelectData(regions: any[]): GenericSelectData[] {
    const selectData: GenericSelectData[] = [];
    if (!Array.isArray(regions) || regions.length === 0) {
        return selectData;
    }
    regions.forEach((region) => {
        const data: GenericSelectData = {};
        data.label = region.name;
        data.value = region.slug;
        selectData.push(data);
    });

    return selectData;
}

// Depends on the regions
export function prepareSizeSelectData(sizes: any[]): GenericSelectData[] {
    const selectData: GenericSelectData[] = [];

    if (!Array.isArray(sizes) || sizes.length === 0) {
        return selectData;
    }
    sizes.forEach((size) => {
        const memory = size.Memory / 1024;
        const data: GenericSelectData = {};
        data.label = `Memory: ${memory}GB Disk: ${size.Disk}GB USD${size.PriceMonthly}`;
        data.value = size;
        selectData.push(data);
    });

    return selectData;
}

export function prepareImageSelectData(images: any[]): GenericSelectData[] {
    const selectData: GenericSelectData[] = [];
    if (!Array.isArray(images) || images.length === 0) {
        return selectData;
    }
    images.forEach((image) => {
        const data: GenericSelectData = {};
        data.label = `${image.name}(${image.distribution})`;
        data.value = image.id;
        selectData.push(data);
    });

    return selectData;
}

export const getStyleForReactSelect = (theme) => {
    if (!theme) {
        return null;
    }

    return {
        menuPortal: (provided) => ({
            ...provided,
            zIndex: 9999,
        }),
        control: (provided, state) => ({
            ...provided,
            color: theme.centerChannelColor,
            background: theme.centerChannelBg,

            // Overwrittes the different states of border
            borderColor: state.isFocused ? changeOpacity(theme.centerChannelColor, 0.25) : changeOpacity(theme.centerChannelColor, 0.2),
            padding: '2px 4px 2px 6px',

            // Removes weird border around container
            boxShadow: 'inset 0 1px 1px ' + changeOpacity(theme.centerChannelColor, 0.075),
            borderRadius: '4px',

            '&:hover': {
                borderColor: changeOpacity(theme.centerChannelColor, 0.25),
            },
        }),
        option: (provided, state) => ({
            ...provided,
            background: state.isFocused ? changeOpacity(theme.centerChannelColor, 0.12) : theme.centerChannelBg,
            color: theme.centerChannelColor,
            '&:hover': {
                background: changeOpacity(theme.centerChannelColor, 0.12),
            },
        }),
        clearIndicator: (provided) => ({
            ...provided,
            width: '34px',
            color: changeOpacity(theme.centerChannelColor, 0.4),
            transform: 'scaleX(1.15)',
            marginRight: '-10px',
            '&:hover': {
                color: theme.centerChannelColor,
            },
        }),
        multiValue: (provided) => ({
            ...provided,
            background: changeOpacity(theme.centerChannelColor, 0.15),
        }),
        multiValueLabel: (provided) => ({
            ...provided,
            color: theme.centerChannelColor,
            paddingBottom: '4px',
            paddingLeft: '8px',
            fontSize: '90%',
        }),
        multiValueRemove: (provided) => ({
            ...provided,
            transform: 'translateX(-2px) scaleX(1.15)',
            color: changeOpacity(theme.centerChannelColor, 0.4),
            '&:hover': {
                background: 'transparent',
            },
        }),
        menu: (provided) => ({
            ...provided,
            color: theme.centerChannelColor,
            background: theme.centerChannelBg,
            border: '1px solid ' + changeOpacity(theme.centerChannelColor, 0.2),
            borderRadius: '0 0 2px 2px',
            boxShadow: changeOpacity(theme.centerChannelColor, 0.2) + ' 1px 3px 12px',
            marginTop: '4px',
        }),
        input: (provided) => ({
            ...provided,
            color: theme.centerChannelColor,
        }),
        placeholder: (provided) => ({
            ...provided,
            color: theme.centerChannelColor,
        }),
        dropdownIndicator: (provided) => ({
            ...provided,

            '&:hover': {
                color: theme.centerChannelColor,
            },
        }),
        singleValue: (provided) => ({
            ...provided,
            color: theme.centerChannelColor,
        }),
        indicatorSeparator: (provided) => ({
            ...provided,
            display: 'none',
        }),
    };
};