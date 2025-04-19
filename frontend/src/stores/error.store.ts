import { defineStore } from "pinia";

export interface ErrorState {
    show: boolean
    message: string
}

export const useErrorStore = defineStore("error", {
    state: (): ErrorState => ({
        show: false,
        message: ""
    }),
    actions: {
        summon(message: string) {
            this.show = true
            this.message = message
        },
        close() {
            this.show = false
            this.message = ""
        }
    }
})