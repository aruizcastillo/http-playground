import { Component, inject, signal } from '@angular/core'
import { FormsModule } from '@angular/forms'
import { MessageService } from './services/message.service'

@Component({
  selector: 'app-root',
  imports: [FormsModule],
  templateUrl: './app.html',
})
export class App {
  private readonly messageService = inject(MessageService)

  readonly getResponse = signal('')
  readonly postMessage = signal('Hello from ANGULAR frontend')
  readonly postRequest = signal('')
  readonly postResponse = signal('')

  handleGet() {
    this.messageService.fetchMessage().subscribe((data) => {
      this.getResponse.set(data.message)
    })
  }

  handlePost() {
    const message = this.postMessage()
    this.postRequest.set(message)

    this.messageService.sendMessage(message).subscribe((data) => {
      this.postResponse.set(data.message)
    })
  }
}
