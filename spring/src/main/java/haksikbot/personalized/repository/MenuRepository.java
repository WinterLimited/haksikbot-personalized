package haksikbot.personalized.repository;

import haksikbot.personalized.domain.Menu;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;

public interface MenuRepository extends JpaRepository<Menu, Long> {
    // MenuName으로 Menu를 찾는다
    Optional<Menu> findByMenuName(String menuName);
}
