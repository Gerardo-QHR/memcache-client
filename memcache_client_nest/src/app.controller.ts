
import { Controller, Get, Post, Param, Query } from '@nestjs/common';
import { MemcachedService } from './memcached/memcached.service';

@Controller()
export class AppController {
  constructor(private readonly memcachedService: MemcachedService) {}

  @Post('/set/:key')
  async setKeyValue(
    @Param('key') key: string,
    @Query('value') value: string,
  ): Promise<any> {
    await this.memcachedService.setValue(key, value, 3600); // Set value with expiration of 1 hour (3600 seconds)
    return { message: `Key '${key}' set successfully in Memcached` };
  }

  @Get('/get/:key')
  async getKeyValue(@Param('key') key: string): Promise<any> {
    const value = await this.memcachedService.getValue(key);
    return { key, value };
  }
}
