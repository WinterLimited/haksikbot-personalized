package haksikbot.personalized.service;

import haksikbot.personalized.domain.Menu;
import haksikbot.personalized.repository.MenuRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Optional;

@Service
@Transactional
@RequiredArgsConstructor
public class MenuService {

    private final MenuRepository menuRepository;

    public Optional<Menu> findOne(String menuName) {
        return menuRepository.findByMenuName(menuName);
    }

    public void save(Menu menu) {
        menuRepository.save(menu);
    }

    /**
     * 메뉴 총 평균 점수를 업데이트한다
     * 메뉴 DB에 없는 메뉴를 입력하면 새로운 메뉴를 추가한다
     * 메뉴 DB에 존재하는 메뉴를 입력하면 점수를 업데이트한다
     */
    @Transactional
    public void updateMenuScore(String menuName, int score) {
        Optional<Menu> menu = menuRepository.findByMenuName(menuName);
        // menu domain에 insertOrUpdate 메서드를 활용
        if (menu.isPresent()) {
            // 값이 존재하면 점수를 업데이트한다
            // updateMenu() 메서드를 활용
            menu.get().updateMenu(score);
        } else {
            // 값이 존재하지 않으면 새로운 메뉴를 추가한다
            menuRepository.save(Menu.createMenu(menuName, score));
        }
    }

}
