package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/whereabouts/web-template-ddd/domain/entity"
	"github.com/whereabouts/web-template-ddd/domain/repository"
)

type TestService struct {
	repo repository.TestRepository
}

func NewTestService() *TestService {
	return &TestService{
		repo: repository.GetTestRepository(),
	}
}

func (service *TestService) TestRegister(ctx context.Context, name string, password string) (*entity.Test, error) {
	// domain 中执行具体的业务逻辑，封装完整的数据，供application可直接使用；
	// 若业务逻辑过于简单（如简单的增删查改），可跳过domain层，在application中直接使用repository进行操作
	test := &entity.Test{
		Name:     name,
		Password: password,
	}
	err := test.MD5Password()
	if err != nil {
		return nil, errors.Wrap(err, "md5 password err")
	}

	err = service.repo.Register(ctx, test)
	if err != nil {
		return nil, errors.Wrap(err, "test repo register err")
	}

	return test, nil
}
