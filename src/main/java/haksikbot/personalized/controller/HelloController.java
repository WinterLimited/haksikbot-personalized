package haksikbot.personalized.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Arrays;
import java.util.List;

@RestController
public class HelloController {

    // react와 연동 테스트용
    @GetMapping("/hello")
    public List<String> hello() {
        return Arrays.asList("hello", "world");
    }

}
