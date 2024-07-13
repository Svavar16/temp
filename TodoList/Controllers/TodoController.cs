using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using TodoList.Data;
using TodoList.Models;

namespace TodoList.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class TodoController : ControllerBase
    {
        private readonly TodoDbContext context;

        public TodoController(TodoDbContext context)
        {
            this.context = context;
        }

        [HttpGet]
        public async Task<ActionResult<List<Todo>>> GetAllTodos()
        {
            return Ok(await context.Todos.ToListAsync<Todo>());
        }

        [HttpGet("{id}")]
        public async Task<ActionResult<Todo>> GetSingleTodo([FromRoute] Guid id)
        {
            return Ok(await context.Todos.FirstOrDefaultAsync(c => c.Id == id));
        }

        [HttpPost]
        public async Task<ActionResult> CreateTodo(Todo item)
        {
            if (ModelState != null)
            {
                return BadRequest(ModelState);
            }

            this.context.Todos.Add(item);
            await this.context.SaveChangesAsync();

            return Ok();
        }

        [HttpPut("{id}")]
        public async Task<ActionResult<Todo>> UpdateTodo([FromBody] Todo itemToUpdate, [FromRoute] Guid id)
        {
            var item = await context.Todos.FirstOrDefaultAsync(c => c.Id == id);
            if (item == null)
            {
                return BadRequest(item);
            }

            item.TodoItem = itemToUpdate.TodoItem;

            this.context.Todos.Update(item);
            await this.context.SaveChangesAsync();
            return Ok(item);
        }

        [HttpDelete("{id}")]
        public async Task<ActionResult> DeleteTodo([FromRoute] Guid id)
        {
            var item = await this.context.Todos.FirstOrDefaultAsync(c => c.Id == id);
            if (item == null)
            {
                return BadRequest(item);
            }

            this.context.Todos.Remove(item);
            await this.context.SaveChangesAsync();
            return NoContent();
        }
    }
}
