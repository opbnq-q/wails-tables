import { toSimpleDateString } from "../../utils/date/converters.util"

type Color = "red" | "yellow" | "blue" | "green" | "pink"

export class CalendarEvent {
    id: string
    title: string
    description: string
    color: Color
    with?: string
    time: { start: string; end: string }
  
    constructor(
      title: string,
      description: string,
      timeStart: Date,
      timeEnd: Date,
      options: { withPerson?: string,
        color: Color }
    ) {
      this.title = title
      this.description = description
      this.time = {
        start: toSimpleDateString(timeStart),
        end: toSimpleDateString(timeEnd),
      }
      this.with = options.withPerson
      this.color = options.color ?? 'red'
      this.id = new Date().getTime().toString()
    }
}