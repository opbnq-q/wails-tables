export const structView = (item: any, path: any) => {
    if (!item) return;
    if (!path?.length) return item;
    let result = item
    let i = 0
    let current
    while (current != path[path.length - 1] && result) {
      current = path[i]
      result = result[current]
      i++
    }
    return result
}

export const manyStructsView = (items: any, path: any) => {
    if (!Array.isArray(items)) return structView(items, path);
    return items.map(item => structView(item, path)).join(", ")
}