import {WoodSpecType} from "../../bindings/wails_tables/internal/models";
import type {IService} from "../types/service.type.ts";
import {
  Count,
  Create,
  Delete,
  GetAll,
  GetById, SearchByAllTextFields, SortedByOrder,
  Update} from '../../bindings/wails_tables/internal/services/woodspectypeservice.ts'
import type {SortOptions} from "../types/sort-options.type.ts";


export default class WoodSpecTypeService implements IService<WoodSpecType> {
  async read(id: number) {
    return (await GetById(id)) as WoodSpecType;
  }

  async readAll() {
    return (await GetAll()) as WoodSpecType[];
  }

  async create(item: WoodSpecType) {
    await Create(item);
  }

  async delete(id: number) {
    return await Delete(id);
  }

  async update(item: WoodSpecType) {
    await Update(item);
  }

  async count() {
    return await Count();
  }

  async search(input: string) {
    return (await SearchByAllTextFields(input)) as WoodSpecType[];
  }

  async sort(options: SortOptions<WoodSpecType>) {
    return (await SortedByOrder(
      Object.entries(options)
        .map((item) => {
          if (item[1] !== "NONE") {
            return {
              Name: item[0],
              Order: item[1],
            };
          }
        })
        .filter((item) => !!item),
    )) as WoodSpecType[];
  }
}
