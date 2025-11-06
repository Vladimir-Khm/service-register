package user

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"service-register/internal/config"
	"service-register/internal/models"
	"service-register/mocks"
)

type UserSuite struct {
	suite.Suite
	mock           sqlmock.Sqlmock
	userService    *UserService
	userRepository *mocks.UserRepository
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, &UserSuite{})
}

func (suite *UserSuite) SetupTest() {
	userRepositoryMock := &mocks.UserRepository{}
	suite.userService = CreateUserService(&config.Config{}, slog.Logger{}, userRepositoryMock)
	suite.userRepository = userRepositoryMock
}

func (suite *UserSuite) TestGetUserReactions_ShouldPass() {
	require := require.New(suite.T())
	userID := int64(1)
	chatID := int64(2)
	reactions := []models.Reaction{
		{ID: 1, AuthorID: userID, MessageID: 1, Emoji: "👍"},
		{ID: 2, AuthorID: userID, MessageID: 2, Emoji: "👎"},
	}

	suite.userRepository.On("GetUserBasicReactions", userID, &chatID).Return(reactions, nil)

	result, err := suite.userService.GetUserReactions(userID, &chatID)
	require.NoError(err)
	require.Len(result, 2)
	require.Equal(result[0].Emoji, "👍")
	require.Equal(result[1].Emoji, "👎")
}

func (suite *UserSuite) TestGetUserReactions_ShouldError() {
	require := require.New(suite.T())
	userID := int64(1)
	chatID := int64(2)

	suite.userRepository.On("GetUserBasicReactions", userID, &chatID).Return(nil, errors.New("error"))

	result, err := suite.userService.GetUserReactions(userID, &chatID)
	require.Nil(result)
	require.Error(err)
}

func (suite *UserSuite) TestGetUserGetDTO_ShouldPass() {
	require := require.New(suite.T())
	userID := int64(1)
	chatID := int64(2)
	user := &models.ShortUser{
		ID:        userID,
		FirstName: "User1",
	}

	suite.userRepository.On("GetShortUserByID", userID, &chatID).Return(user, nil)

	result, err := suite.userService.GetUserGetDTO(userID, &chatID)
	require.NoError(err)
	require.Equal(result.FirstName, "User1")
}

func (suite *UserSuite) TestGetUserGetDTO_ShouldError() {
	require := require.New(suite.T())
	userID := int64(1)
	chatID := int64(2)

	suite.userRepository.On("GetShortUserByID", userID, &chatID).Return(nil, errors.New("error"))

	result, err := suite.userService.GetUserGetDTO(userID, &chatID)
	require.Nil(result)
	require.Error(err)
}

//TODO checkInitData should pass

func (suite *UserSuite) TestCheckInitData_ShouldError() {
	require := require.New(suite.T())
	initData := "invalid_init_data"

	resultUser, resultToken, err := suite.userService.CheckInitData(initData)
	require.Error(err)
	require.Nil(resultUser)
	require.Empty(resultToken)
}

func (suite *UserSuite) TestAddUserIfNotExists_ShouldPass() {
	require := require.New(suite.T())
	initData := &initdata.InitData{
		User: initdata.User{
			ID:        1,
			FirstName: "User1",
			Username:  "user1",
			LastName:  "Last1",
		},
	}
	user := &models.User{
		ID:        1,
		FirstName: "User1",
		Username:  &initData.User.Username,
		LastName:  &initData.User.LastName,
	}

	suite.userRepository.On("CreateOrUpdateUser", user).Return(nil)

	result, err := suite.userService.AddUserIfNotExists(initData)
	require.NoError(err)
	require.Equal(result.FirstName, "User1")
}

func (suite *UserSuite) TestAddUserIfNotExists_ShouldError() {
	require := require.New(suite.T())
	initData := &initdata.InitData{
		User: initdata.User{
			ID:        1,
			FirstName: "User1",
			Username:  "user1",
			LastName:  "Last1",
		},
	}
	user := &models.User{
		ID:        1,
		FirstName: "User1",
		Username:  &initData.User.Username,
		LastName:  &initData.User.LastName,
	}

	suite.userRepository.On("CreateOrUpdateUser", user).Return(errors.New("error"))

	result, err := suite.userService.AddUserIfNotExists(initData)
	require.Error(err)
	require.Nil(result)
}

func (suite *UserSuite) TestCreateOrUpdateUser_ShouldPass() {
	require := require.New(suite.T())
	user := &models.User{
		ID:        1,
		FirstName: "User1",
		Username:  nil,
		LastName:  nil,
	}

	suite.userRepository.On("CreateOrUpdateUser", user).Return(nil)

	result, err := suite.userService.CreateOrUpdateUser(user)
	require.NoError(err)
	require.Equal(result.FirstName, "User1")
}

func (suite *UserSuite) TestCreateOrUpdateUser_ShouldError() {
	require := require.New(suite.T())
	user := &models.User{
		ID:        1,
		FirstName: "User1",
		Username:  nil,
		LastName:  nil,
	}

	suite.userRepository.On("CreateOrUpdateUser", user).Return(errors.New("error"))

	result, err := suite.userService.CreateOrUpdateUser(user)
	require.Error(err)
	require.Nil(result)
}

//TODO VerifyUser_ShouldPass

func (suite *UserSuite) TestVerifyUser_ShouldError() {
	require := require.New(suite.T())
	userID := int64(1)
	wallet := "valid_wallet"

	suite.userRepository.On("VerifyUser", userID, wallet).Return(errors.New("error"))
	err := suite.userService.VerifyUser(userID, wallet)
	require.Error(err)
}

func (suite *UserSuite) TestVerifyUser_ShouldErrorSecond() {
	require := require.New(suite.T())
	userID := int64(1)
	wallet := "valid_wallet"

	suite.userRepository.On("VerifyUser", userID, wallet).Return(nil)
	suite.userRepository.On("UpdateUserReactionWeight", userID, 1).Return(errors.New("error"))
	err := suite.userService.VerifyUser(userID, wallet)
	require.Error(err)
}
