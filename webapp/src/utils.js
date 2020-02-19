export function prepareRegionsSelectData(regions) {
    const selectData = [];
    regions.forEach((region) => {
        const data = {};
        data.label = region.name;
        data.value = region.slug;
        selectData.push(data);
    });

    return selectData;
}

export function prepareSizeSelectData(sizes) {
    const selectData = [];
    sizes.forEach((size) => {
        const data = {};
        data.label = size.memory;
        data.value = size.slug;
        selectData.push(data);
    });

    return selectData;
}

export function prepareImageSelectData(images) {
    const selectData = [];
    images.forEach((image) => {
        const data = {};
        data.label = image.distribution;
        data.value = image.id;
        selectData.push(data);
    });

    return selectData;
}