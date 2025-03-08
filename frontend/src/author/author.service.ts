import { GetAll, Create, Delete, GetById, Update, Count } from "../../bindings/app/internal/services/authorservice.ts"
import type { Author } from "../../bindings/app/internal/services/models.ts"
import type { IService } from "../types/service.type.ts"
	
export default class AuthorService implements IService<Author> {
	async read(id: number) {
		return await GetById(id) as Author
	}

	async readAll() {
		return await GetAll() as Author[]
	}

	async create(item: Author) {
		await Create(item)
	}

	async delete(id: number) {
		return await Delete(id)
	}

	async update(item: Author) {
		await Update(item)
	}

	async count() {
		return await Count()
	}
}
