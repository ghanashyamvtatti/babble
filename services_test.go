package main
import (
   "testing"
   "ds-project/services"
   "ds-project/config"
   "sync"
   "log"
)


func TestUserNameExists(t *testing.T){
	log.Println("Testing user name exists")
	appConfig := config.NewAppConfig()
	exists := services.CheckUserNameExists(appConfig, "varun")
	if !exists {
		t.Error("fails")
	}
}


func TestTokenValid(t *testing.T){
	log.Println("Testing user name exists")
	appConfig := config.NewAppConfig()
	token := services.GenerateAccessToken(appConfig, "varun")
	if !services.CheckAccessTokenValid(appConfig,"varun",token) {
		t.Error("Fails")
	}
	log.Println("token valid")
}

func TestGetPostsForUsers(t *testing.T){
	log.Println("Testing get post service")
	appConfig := config.NewAppConfig()
	posts := services.GetPostsForUser(appConfig, "varun")

	log.Println(posts[0].Post)
	
	if posts[0].Post != "My name is Varun."{
		t.Error("fails")
	}
}

func TestPostAdd(t *testing.T){
	log.Println("Testing add post service")
	appConfig := config.NewAppConfig()
	services.AddPost(appConfig, "varun", "New POST")

	posts := services.GetPostsForUser(appConfig, "varun")
	size := len(posts)
	log.Println(posts[size-1].Post)
	
	if posts[size-1].Post != "New POST"{
		t.Error("fails")
	}
}


func TestMultiplePost(t *testing.T) {
	log.Println("Testing add multiple post service")

	appConfig := config.NewAppConfig()
	initialPosts := services.GetPostsForUser(appConfig, "varun")
	initialPostsLength := len(initialPosts)

	wg := sync.WaitGroup{}
	for idx := 0; idx < 1000; idx++ {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()
			p := "New POST " + string(idx)
			services.AddPost(appConfig, "varun", p)
		}(idx)
	}

	wg.Wait()
	finalPosts := services.GetPostsForUser(appConfig, "varun")
	finalPostsLength := len(finalPosts)
	
	log.Println(initialPostsLength)
	log.Println(finalPostsLength)

	if finalPostsLength != initialPostsLength + 1000 {
		t.Error("fails")
	}

	// for i := 0; i < finalPostsLength; i++ {
	// 	log.Println(finalPosts[i].Post)
	// }
}


func TestGetFeedForUsers(t *testing.T){
	log.Println("Testing get post service")
	appConfig := config.NewAppConfig()
	feeds := services.GetFeedForUsername(appConfig, "varun")

	log.Println(feeds[0].Post)
	
	if len(feeds) == 0 {
		t.Error("fails")
	}
}