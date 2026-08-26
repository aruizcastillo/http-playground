import { inject, Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'

@Injectable({
  providedIn: 'root',
})
export class MessageService {
  private readonly http = inject(HttpClient)
  private readonly apiUrl = 'http://localhost:3000'

  fetchMessage() {
    return this.http.get<{ message: string }>(
      `${this.apiUrl}/api/message`,
    )
  }

  sendMessage(message: string) {
    return this.http.post<{ received: string }>(
      `${this.apiUrl}/api/message`,
      { message },
    )
  }
}