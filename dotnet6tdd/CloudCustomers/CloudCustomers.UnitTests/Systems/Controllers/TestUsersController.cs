using CloudCustomers.API.Controllers;
using FluentAssertions;
using Microsoft.AspNetCore.Mvc;
using Moq;
using Xunit;

namespace CloudCustomers.UnitTests.Systems.Controller;

public class TestUsersContoller
{
    [Fact]
    public async Task Get_OnSuccess_ReturnStatusCode200()
    {
        // Arrange
        var mockUserService = new Mock<IUserService>();
        var sut = new UsersController(mockUserService.Object);
        // Act
        var result = (OkObjectResult)await sut.Get();

        // Assert
        result.StatusCode.Should().Be(200);
    }

    [Fact]
    public async Task Get_On_Success_InvokeUserService()
    {
        var mockUserService = new Mock<IUserService>();
        var sut = new UsersController(mockUserService.Object);

        var result = (OkObjectResult)await sut.Get();
    }
}