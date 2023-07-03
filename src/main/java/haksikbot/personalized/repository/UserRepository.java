package haksikbot.personalized.repository;

import haksikbot.personalized.domain.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository  extends JpaRepository<User, Long> {
}
