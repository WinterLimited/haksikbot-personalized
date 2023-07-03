package haksikbot.personalized.controller;

import haksikbot.personalized.service.MenuService;
import haksikbot.personalized.service.UserMenuScoreService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequiredArgsConstructor
public class UserMenuScoreController {

    private final MenuService menuService;
    private final UserMenuScoreService userMenuScoreService;

    // rest api 방식으로 GET 요청을 받는다
    // 요청을 받으면 사용자의 메뉴 점수의 적합도를 계산해서 반환한다
    // menuName의 예시 형식 '닭갈비야채덮밥, 핫도그, 단무지, 열무김치, 찜두부'
    // menuName은 ','로 split해서 배열로 만들어서 JPA 메서드로 넘겨줘야함
    // menuName의 배열을 for문으로 돌려서 메뉴가 존재하는지 확인하고
    // 존재하면 점수를 가져오고
    // 존재하지 않으면 전체 사용자의 평균 값인 해당 Menu의 score를 가져온다
    @GetMapping("/api/{userId}/{menuName}")
    public int getUserMenuScore(@RequestParam Long userId, @RequestParam String menuName) {
        menuName = menuName.replaceAll(" ", "");
        String[] menuNameArray = menuName.split(",");
        int scoreSum = 0;
        int scoreCount = 0;
        for (String name : menuNameArray) {
            if(userMenuScoreService.findByUserIdAndMenuMenuName(userId, name).isPresent()) {
                scoreSum += userMenuScoreService.findByUserIdAndMenuMenuName(userId, name).get().getScore();
                scoreCount++;
            }
            else {
                scoreSum += menuService.findOne(name).get().getScore();
                scoreCount++;
            }
        }
        return scoreSum / scoreCount;
    }

    // rest api 방식으로 POST 요청을 받는다
    // 요청을 받으면 사용자의 메뉴 점수를 업데이트한다
    // menuName의 예시 형식 '닭갈비야채덮밥, 핫도그, 단무지, 열무김치, 찜두부'
    // menuName은 ','로 split해서 배열로 만들어서 JPA 메서드로 넘겨줘야함
    // menuName의 배열을 for문으로 돌려서 메뉴가 존재하는지 확인하고
    // 존재하면 점수를 업데이트하고
    // 존재하지 않으면 새로운 메뉴를 추가한다
    @PostMapping("/api/menuscore")
    public void updateUserMenuScore(@RequestParam Long userId, @RequestParam String menuName, @RequestParam int score) {
        menuName = menuName.replaceAll(" ", "");
        String[] menuNameArray = menuName.split(",");
        for (String name : menuNameArray) {
            menuService.updateMenuScore(name, score);
            userMenuScoreService.updateMenuScore(userId, name, score);
        }
    }
}
