import { Injectable } from '@nestjs/common';

// ES7 decorators
@Injectable()
export class AppService {
  getHello(): string {
    return 'Hello World!';
  }
}
