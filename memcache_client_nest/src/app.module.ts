import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MemcachedService } from './memcached/memcached.service';

@Module({
  imports: [],
  controllers: [AppController],
  providers: [AppService, MemcachedService],
})
export class AppModule {}
