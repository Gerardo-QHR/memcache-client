package com.example.memcache_client_spring.service;

import net.spy.memcached.MemcachedClient;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.Future;

@Service
public class MemcachedService {
    private final MemcachedClient memcachedClient;

    public MemcachedService() throws IOException {
        memcachedClient = new MemcachedClient(new InetSocketAddress("memcached", 11211));
    }

    public void setValue(String key, Object value, int expiration) {
        try {
            Future<Boolean> future = memcachedClient.set(key, expiration, value);
            if (Boolean.FALSE.equals(future.get())) {
                throw new RuntimeException("Failed to set key '" + key + "' in Memcached");
            }
        } catch (InterruptedException | ExecutionException e) {
            throw new RuntimeException("Error setting key '" + key + "' in Memcached", e);
        }
    }

    public Object getValue(String key) {
        return memcachedClient.get(key);
    }
}
