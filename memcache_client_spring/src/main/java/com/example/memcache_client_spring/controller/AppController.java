package com.example.memcache_client_spring.controller;

import com.example.memcache_client_spring.service.MemcachedService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class AppController {
    @Autowired
    private MemcachedService memcachedService;

    @PostMapping("/set/{key}")
    public String setKeyValue(@PathVariable String key, @RequestParam String value) {
        memcachedService.setValue(key, value, 3600);
        return String.format("Key '%s' set successfully in Memcached", key);
    }

    @GetMapping("/get/{key}")
    public Object getKeyValue(@PathVariable String key) {
        return memcachedService.getValue(key);
    }
}
