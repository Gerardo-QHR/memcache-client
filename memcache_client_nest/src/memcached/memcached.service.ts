import { Injectable } from '@nestjs/common';
import * as Memcached from 'memcached';

@Injectable()
export class MemcachedService {
  private readonly client: Memcached;

  constructor() {
    this.client = new Memcached('memcached:11211'); 
  }

  async setValue(key: string, value: any, expiration: number): Promise<void> {
    return new Promise<void>((resolve, reject) => {
      this.client.set(key, value, expiration, (err) => {
        if (err) {
          reject(err);
        } else {
          resolve();
        }
      });
    });
  }

  async getValue(key: string): Promise<any> {
    return new Promise<any>((resolve, reject) => {
      this.client.get(key, (err, data) => {
        if (err) {
          reject(err);
        } else {
          resolve(data);
        }
      });
    });
  }
}
