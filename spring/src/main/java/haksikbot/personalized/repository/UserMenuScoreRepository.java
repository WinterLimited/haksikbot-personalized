package haksikbot.personalized.repository;

import haksikbot.personalized.domain.UserMenuScore;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;

public interface UserMenuScoreRepository extends JpaRepository<UserMenuScore, Long> {
    Optional<UserMenuScore> findByUserIdAndMenuMenuName(Long userId, String menuName);

}
