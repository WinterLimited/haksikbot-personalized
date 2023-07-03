package haksikbot.personalized.service;

import haksikbot.personalized.domain.Menu;
import haksikbot.personalized.domain.User;
import haksikbot.personalized.domain.UserMenuScore;
import haksikbot.personalized.repository.MenuRepository;
import haksikbot.personalized.repository.UserMenuScoreRepository;
import haksikbot.personalized.repository.UserRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Optional;

@Service
@Transactional
@RequiredArgsConstructor
public class UserMenuScoreService {

    private final UserRepository userRepository;
    private final MenuRepository menuRepository;
    private final UserMenuScoreRepository userMenuScoreRepository;

    public Optional<UserMenuScore> findOne(Long id) {
        return userMenuScoreRepository.findById(id);
    }

    public void save(UserMenuScore userMenuScore) {
        userMenuScoreRepository.save(userMenuScore);
    }

    /**
     * 사용자의 메뉴 점수를 업데이트한다
     * 메뉴 DB에 없는 메뉴를 입력하면 새로운 메뉴를 추가한다
     * 메뉴 DB에 존재하는 메뉴를 입력하면 점수를 업데이트한다
     */
    @Transactional
    public void updateMenuScore(Long userId, String menuName, int score) {
        Optional<UserMenuScore> userMenuScore = userMenuScoreRepository.findByUserIdAndMenuMenuName(userId, menuName);
        if (userMenuScore.isPresent()) {
            // 값이 존재하면 점수를 업데이트한다
            // updateMenu() 메서드를 활용
            userMenuScore.get().updateMenu(score);
        } else {
            // 값이 존재하지 않으면 새로운 메뉴를 추가한다
            // 생성메서드로 user, menu(score 정리)된 객체를 생성메서드로 넘겨줘야함
            User user = userRepository.findById(userId).get();
            Menu menu = menuRepository.findByMenuName(menuName).get();
            userMenuScoreRepository.save(UserMenuScore.createUserMenuScore(user, menu, score));
        }
    }
}
