import { Component, inject } from '@angular/core'
import { FormsModule } from '@angular/forms'
import { MessageService } from './message.service'

@Component({
  selector: 'app-root',
  imports: [FormsModule],
  templateUrl: './app.html',
})
export class App {
  private readonly messageService = inject(MessageService)

  getResponse = ''
  postMessage = ''
  postResponse = ''

  handleGet() {
    this.messageService.fetchMessage().subscribe((data) => {
      this.getResponse = data.message
    })
  }

  handlePost() {
    this.messageService.sendMessage(this.postMessage).subscribe((data) => {
      this.postResponse = data.received
    })
  }
}
