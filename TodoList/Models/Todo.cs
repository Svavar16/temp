namespace TodoList.Models
{
    public class Todo
    {
        public Guid Id { get; set; }
        public string TodoItem { get; set; } = string.Empty;
        // we will ad the user_id later on, maybe?
    }
}
